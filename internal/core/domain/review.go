package domain

type ReviewModel struct {
	ReviewID  int64
	UserID    int64
	ProductID int64
	Rating    int
	Comment   string
	CreatedAt string
	UpdatedAt string
}

type Review struct {
	ReviewModel
	UserName string
}
