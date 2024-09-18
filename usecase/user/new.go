package user

type repository interface{}

type UseCase struct {
	rep repository
}
