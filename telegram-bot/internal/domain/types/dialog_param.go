package types

const (
	DialogParamCardQuestion = "card.question"
	DialogParamCardFileType = "card.file_type"
	DialogParamCardFileId   = "card.file_id"
	DialogParamCardAnswer   = "card.answer"
)

type DialogParam string

type DialogParams map[DialogParam]string
