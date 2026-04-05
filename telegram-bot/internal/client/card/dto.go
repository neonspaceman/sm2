package card

type CreateCardDto struct {
	UserId   string
	Question string
	Answer   string
	FileType FileType
	FileId   string
}

type Card struct {
	Id string
}
