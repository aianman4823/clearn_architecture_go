package interfaces

import "github.com/aianman4823/clearn_architecture_go/internal/entities"

type TestInterface interface {
	ListTest() ([]*entities.Test, error)
	GetTest(id string) (*entities.Test, error)
}
