package requests

type AppHeaderAuthorization struct {
	Authorization string `header:"Authorization,required"`
}
type HeaderToken struct {
	Authorization string `header:"Authorization" binding:"required,min=20"`
}
