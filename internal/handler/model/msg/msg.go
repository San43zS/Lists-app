package msg

type MSG struct {
	Type string `json:"type"`

	Content Data
}

type Data struct {
	Data []byte `json:"data"`
}

type Test struct {
	Data string `json:"data"`
}
