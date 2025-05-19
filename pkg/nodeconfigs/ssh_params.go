package nodeconfigs

func DefaultSSHParams() *SSHParams {
	return &SSHParams{Port: 22}
}

type SSHParams struct {
	Port int `json:"port"`
}
