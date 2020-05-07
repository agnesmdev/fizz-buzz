package interfaces

type IFizzBuzzService interface {
	GetFizzBuzz(int1 int, int2 int, limit int, str1 string, str2 string) ([]string, error)
}