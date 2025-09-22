package web

type CategoryCreateRequestUpdate struct {
	Id   int    `validate:"required"`
	Name string `validate:"required"`
}
