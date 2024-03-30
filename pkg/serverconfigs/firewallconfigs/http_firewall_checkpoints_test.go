package firewallconfigs_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
)

func TestRuleCheckpoint_Markdown(t *testing.T) {
	var result = []string{}
	for _, def := range firewallconfigs.AllCheckpoints {
		def.Description = strings.ReplaceAll(def.Description, "<code-label>", "`")
		def.Description = strings.ReplaceAll(def.Description, "</code-label>", "`")

		var row = "## " + def.Name + "\n"
		row += "* 名称：" + def.Name + "\n"
		row += "* 代号：`${" + def.Prefix + "}`\n"
		row += "* 描述：" + def.Description + "\n"
		if len(def.Version) > 0 {
			row += "* 版本：v" + def.Version + "\n"
		}
		result = append(result, row)
	}

	fmt.Print(strings.Join(result, "\n") + "\n")
}
