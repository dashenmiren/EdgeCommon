// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://cdn.foyeseo.com .

package nodeconfigs

func DefaultSSHParams() *SSHParams {
	return &SSHParams{Port: 22}
}

type SSHParams struct {
	Port int `json:"port"`
}
