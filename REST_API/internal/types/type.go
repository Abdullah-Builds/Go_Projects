package types

type Student struct{
	Name string `validate:"required"`
	Email string `validate:"required"`
	Age int64  `validate:"required"`
}