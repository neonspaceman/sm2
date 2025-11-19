package card

type CreateRequestDto struct {
	UserId   string
	Question string
	Answer   string
	FileType FileType
	FileId   string
}
