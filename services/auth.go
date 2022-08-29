package services

type LoginService interface {
	LoginUser(email string, password string) bool
}
