package langs

type MessageCode string

func (this MessageCode) For(langCode LangCode, args ...any) string {
	return Message(langCode, this, args...)
}

func (this MessageCode) String() string {
	return string(this)
}
