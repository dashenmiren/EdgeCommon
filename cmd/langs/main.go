// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://cdn.foyeseo.com .

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dashenmiren/EdgeCommon/pkg/langs"
	"github.com/iwind/TeaGo/Tea"
	_ "github.com/iwind/TeaGo/bootstrap"
	"github.com/iwind/TeaGo/lists"
	"github.com/iwind/TeaGo/types"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var args = os.Args
	if len(args) >= 2 {
		switch args[1] {
		case "generate":
			// generate go codes from json files
			runGenerate()
		case "search":
			// search hans from dir path
			runSearch()
		}
	} else {
		fmt.Println("Usage: langs [generate|search]")
	}
}

func runGenerate() {
	var rootDir = filepath.Clean(Tea.Root + "/../pkg/langs/protos")
	dir, err := os.Open(rootDir)
	if err != nil {
		fmt.Println("[ERROR]read dir failed: " + err.Error())
		return
	}
	defer func() {
		_ = dir.Close()
	}()

	files, err := dir.Readdir(0)
	if err != nil {
		fmt.Println("[ERROR]read dir failed: " + err.Error())
		return
	}

	var dirRegexp = regexp.MustCompile(`^[a-z]+-[a-z]+$`)
	var jsonFileNameRegexp = regexp.MustCompile(`^([a-zA-Z0-9]+)(_([a-zA-Z0-9]+))*\.json$`)
	var messageCodeRegexp = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	var jsonCommentRegexp = regexp.MustCompile(`//\s+.+`)

	var messageCodes = []string{}
	var langMaps = map[string]*langs.Lang{} // lang => *langs.Lang
	var defaultLang = langs.DefaultManager().DefaultLang()

	const maxMessageCodeLen = 128

	for _, file := range files {
		var dirName = file.Name()

		if !file.IsDir() || !dirRegexp.MatchString(dirName) {
			continue
		}
		var langCode = dirName
		var isBaseLang = langCode == defaultLang

		var processOk = func() bool {
			jsonFiles, err := filepath.Glob(rootDir + "/" + dirName + "/*.json")
			if err != nil {
				fmt.Println("[ERROR]list json files failed: " + err.Error())
				return false
			}

			for _, jsonFile := range jsonFiles {
				var jsonFileName = filepath.Base(jsonFile)
				if len(jsonFileName) == 0 || !jsonFileNameRegexp.MatchString(jsonFileName) {
					continue
				}
				var module = strings.TrimSuffix(jsonFileName, ".json")

				data, err := os.ReadFile(jsonFile)
				if err != nil {
					fmt.Println("[ERROR]read json file '" + jsonFile + "' failed: " + err.Error())
					return false
				}

				// remove comments in json
				data = jsonCommentRegexp.ReplaceAll(data, []byte{})

				var m = map[string]string{} // code => value
				err = json.Unmarshal(data, &m)
				if err != nil {
					fmt.Println("[ERROR]decode json file '" + jsonFile + "' failed: " + err.Error())
					return false
				}

				var newM = map[string]string{}
				for code, value := range m {
					if !messageCodeRegexp.MatchString(code) {
						fmt.Println("[ERROR]invalid message code '" + code + "'")
						return false
					}

					var fullCode = module + "@" + code

					if len(fullCode) > maxMessageCodeLen {
						fmt.Println("[ERROR]message code '" + fullCode + "' too long, max length: " + types.String(maxMessageCodeLen))
						return false
					}

					if isBaseLang {
						messageCodes = append(messageCodes, fullCode)
					}
					newM[fullCode] = value
				}

				finalLang, ok := langMaps[langCode]
				if !ok {
					finalLang = langs.NewLang(langCode)
					langMaps[langCode] = finalLang
				}
				for code, value := range newM {
					if finalLang.Has(langs.MessageCode(code)) {
						fmt.Println("[ERROR]message code '" + code + "' duplicated")
						return false
					}
					finalLang.Set(langs.MessageCode(code), value)
				}
			}

			return true
		}()
		if !processOk {
			return
		}
	}

	// compile
	for langCode, lang := range langMaps {
		err = lang.Compile()
		if err != nil {
			fmt.Println("[ERROR]compile '" + langCode + "' failed: " + err.Error())
			return
		}
	}

	// check message codes
	fmt.Println("checking message codes ...")
	var defaultMessageMap = map[langs.MessageCode]string{}
	for langCode, messageLang := range langMaps {
		if langCode == defaultLang { // only check lang not equal to 'en-us'
			defaultMessageMap = messageLang.GetAll()
			continue
		}

		for messageCode := range messageLang.GetAll() {
			if !lists.ContainsString(messageCodes, messageCode.String()) {
				fmt.Println("[ERROR]message code '" + messageCode.String() + "' in lang '" + langCode + "' not exist in default lang file ('" + defaultLang + "')")
				return
			}
		}
	}
	fmt.Println("found '" + types.String(len(messageCodes)) + "' message codes")

	// generate codes/codes.go
	sort.Strings(messageCodes)
	var codesSource = `
// generated by run 'langs generate'

package codes

import(
	"github.com/dashenmiren/EdgeCommon/pkg/langs"
)

const (
	`

	for index, messageCode := range messageCodes {
		// add comment to message code
		comment, _, _ := strings.Cut(defaultMessageMap[langs.MessageCode(messageCode)], "\n")

		codesSource += upperWords(messageCode) + " langs.MessageCode = " + strconv.Quote(messageCode) + " // " + comment

		// add NL
		if index != len(messageCodes)-1 {
			codesSource += "\n"
		}
	}

	codesSource += `
)
`

	formattedCodesSource, err := format.Source([]byte(codesSource))
	if err != nil {
		fmt.Println("[ERROR]format 'codes.go' failed: " + err.Error())
		return
	}

	fmt.Println("generating 'codes/codes.go' ...")
	err = os.WriteFile(Tea.Root+"/../pkg/langs/codes/codes.go", formattedCodesSource, 0666)
	if err != nil {
		fmt.Println("[ERROR]write to 'codes.go' failed: " + err.Error())
		return
	}

	// generate messages_LANG.go
	for langCode, messageLang := range langMaps {
		var langFile = strings.ReplaceAll("messages/messages_"+langCode+".go", "-", "_")

		fmt.Println("generating '" + langFile + "' ...")
		var source = `
// generated by run 'langs generate'

package messages

import(
	"github.com/dashenmiren/EdgeCommon/pkg/langs"
)

func init() {
	langs.Load("` + langCode + `", map[langs.MessageCode]string{
	`

		for _, code := range messageCodes {
			var value = messageLang.Get(langs.MessageCode(code))
			source += strconv.Quote(code) + ": " + strconv.Quote(value) + ",\n"
		}

		source += `
	})
}
`

		formattedSource, err := format.Source([]byte(source))
		if err != nil {
			fmt.Println("[ERROR]format '" + langFile + "' failed: " + err.Error())
			return
		}

		err = os.WriteFile(Tea.Root+"/../pkg/langs/"+langFile, formattedSource, 0666)
		if err != nil {
			fmt.Println("[ERROR]write file '" + langFile + "' failed: " + err.Error())
			return
		}
	}

	// generate language javascript files for EdgeAdmin and EdgeUser (commercial versions)
	for lang, messageLang := range langMaps {
		if lang != defaultLang {
			// TODO merge messageMap with default message map
		}

		for _, component := range []string{"EdgeAdmin", "EdgeUser"} {
			fmt.Println("generating '" + lang + ".js' for " + component + " ...")
			var targetFile = filepath.Clean(Tea.Root + "/../../" + component + "/web/public/js/langs/" + lang + ".js")
			var targetDir = filepath.Dir(targetFile)
			dirStat, _ := os.Stat(targetDir)
			if dirStat != nil {
				var prefix = ""
				switch component {
				case "EdgeAdmin":
					prefix = "admin_"
				case "EdgeUser":
					prefix = "user_"
				}
				if len(prefix) == 0 {
					continue
				}

				var filteredMessages = map[langs.MessageCode]string{}
				for code, value := range messageLang.GetAll() {
					if strings.HasPrefix(code.String(), prefix) && strings.Contains(code.String(), "@ui_") /** must contains 'ui' **/ {
						filteredMessages[code] = value
					}
				}

				messageMapJSON, jsonErr := json.Marshal(filteredMessages)
				if jsonErr != nil {
					fmt.Println("[ERROR]marshal message map failed: " + jsonErr.Error())
					return
				}
				err = os.WriteFile(targetFile, []byte(`// generated by 'langs generate'
window.LANG_MESSAGES = `+string(messageMapJSON)+";\n"), 0666)
				if err != nil {
					fmt.Println("[ERROR]write file '" + targetFile + "' failed: " + err.Error())
					return
				}

				// base.js
				if lang == "zh-cn" {
					var baseJSFile = filepath.Dir(targetFile) + "/base.js"
					err = os.WriteFile(baseJSFile, []byte(`// generated by 'langs generate'
window.LANG_MESSAGES_BASE = `+string(messageMapJSON)+";\n"), 0666)
					if err != nil {
						fmt.Println("[ERROR]write file '" + baseJSFile + "' failed: " + err.Error())
						return
					}
				}
			}
		}
	}

	fmt.Println("success")
}

func upperWords(s string) string {
	var pieces = strings.Split(s, "@")
	for pieceIndex, piece := range pieces {
		var words = strings.Split(piece, "_")
		for index, word := range words {
			words[index] = upperWord(word)
		}
		pieces[pieceIndex] = strings.Join(words, "")
	}
	return strings.Join(pieces, "_")
}

func upperWord(word string) string {
	var l = len(word)
	if l == 0 {
		return ""
	}

	if l == 1 {
		return strings.ToUpper(word)
	}

	// special words
	switch word {
	case "api", "http", "https", "tcp", "tls", "ssl", "udp", "ip", "dns", "ns",
		"waf", "acme", "ssh", "toa", "http2", "http3", "uam", "cc",
		"db", "isp", "sni", "ui", "soa", "ocsp", "en", "zh", "ad", "tsig",
		"rpc", "dao":
		return strings.ToUpper(word)
	case "ipv6":
		return "IPv6"
	case "ddos":
		return "DDoS"
	case "webp":
		return "WebP"
	case "doh":
		return "DoH"
	}

	return strings.ToUpper(word[:1]) + word[1:]
}

func runSearch() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: langs search DIR")
		return
	}
	var dir = os.Args[2]
	stat, err := os.Stat(dir)
	if err != nil {
		fmt.Println("[ERROR]could not find dir '" + dir + "': " + err.Error())
		return
	}
	if !stat.IsDir() {
		fmt.Println("[ERROR]could not find dir '" + dir + "'")
		return
	}

	fmt.Println("searching '" + dir + "' ...")

	var ext = ".go"

	var resultFiles = []string{}
	for _, pattern := range []string{
		"*" + ext,
		strings.Repeat("*/", 1) + "*" + ext,
		strings.Repeat("*/", 2) + "*" + ext,
		strings.Repeat("*/", 3) + "*" + ext,
		strings.Repeat("*/", 4) + "*" + ext,
		strings.Repeat("*/", 5) + "*" + ext,
		strings.Repeat("*/", 6) + "*" + ext,
		strings.Repeat("*/", 7) + "*" + ext,
		strings.Repeat("*/", 8) + "*" + ext,
		strings.Repeat("*/", 9) + "*" + ext,
		strings.Repeat("*/", 10) + "*" + ext,
	} {
		goFiles, err := filepath.Glob(dir + "/" + pattern)
		if err != nil {
			fmt.Println("[ERROR]search error: " + err.Error())
			return
		}
		resultFiles = append(resultFiles, goFiles...)
	}

	if len(resultFiles) == 0 {
		fmt.Println("no files found in the dir")
		return
	}

	if err != nil {
		fmt.Println("[ERROR]search dir '" + dir + "' failed: " + err.Error())
		return
	}
	var hansRegexp = regexp.MustCompile(`\p{Han}+`)
	var countMatches = 0
	for _, goFile := range resultFiles {
		if strings.HasSuffix(goFile, "_test.go") ||
			strings.HasSuffix(goFile, "_plus_test.go") ||
			strings.Contains(goFile, "/messages/messages_") {
			continue
		}

		data, err := os.ReadFile(goFile)
		if err != nil {
			fmt.Println("[ERROR]read file '" + goFile + "' failed: " + err.Error())
			return
		}
		var matches = hansRegexp.FindAllSubmatchIndex(data, -1)
		if len(matches) > 0 {
			for _, match := range matches {
				// ignore comment
				switch ext {
				case ".go":
					if checkIsInGoComment(data, match[0]) {
						continue
					}
				}

				countMatches++

				fmt.Printf("%s %s\n", goFile+":"+types.String(bytes.Count(data[:match[0]], []byte{'\n'})+1), string(data[match[0]:match[1]]))
			}
		}
	}
	fmt.Println(countMatches, "matches")
}

func checkIsInGoComment(data []byte, start int) bool {
	if start <= 1 {
		return false
	}

	for {
		start--
		if start <= 1 || data[start] == '\n' {
			return false
		}

		// 'SPACE //'
		if data[start] == '/' && data[start-1] == '/' {
			return true
		}

		// '/** SOMETHING **/'
		if data[start] == '*' && data[start-1] == '/' {
			return true
		}
	}
}
