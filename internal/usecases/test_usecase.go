package usecases

import (
	"github.com/aianman4823/clearn_architecture_go/internal/entities"
	"github.com/aianman4823/clearn_architecture_go/internal/interfaces"
)

type TestUsecase struct {
	repository interfaces.TestInterface
}

func NewTestUsecase(r interfaces.TestInterface) *TestUsecase {
	return &TestUsecase{
		repository: r,
	}
}

func (u *TestUsecase) ListTest() ([]*entities.Test, error) {
	return u.repository.ListTest()
}

func (u *TestUsecase) GetTest(id string) (*entities.Test, error) {
	return u.repository.GetTest(id)
}
