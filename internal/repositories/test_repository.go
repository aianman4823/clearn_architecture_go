package repositories

import (
	"errors"

	"github.com/aianman4823/clearn_architecture_go/internal/entities"
	"gorm.io/gorm"
)

type TestRepository struct {
	Conn *gorm.DB
}

type Test struct {
	Id   int
	Name string
}

func NewTestRepository(conn *gorm.DB) *TestRepository {
	return &TestRepository{Conn: conn}
}

func (r *TestRepository) ListTest() ([]*entities.Test, error) {
	var objs []Test
	result := r.Conn.Table("test").Find(&objs)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	var returnObjs []*entities.Test
	for _, obj := range objs {
		returnObjs = append(returnObjs, convertTestRepositoryModelToEntity(&obj))
	}
	return returnObjs, nil
}

func (r *TestRepository) GetTest(id string) (*entities.Test, error) {
	var obj Test
	result := r.Conn.Table("test").Where("id = ?", id).First(&obj)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return convertTestRepositoryModelToEntity(&obj), nil
}

func convertTestRepositoryModelToEntity(obj *Test) *entities.Test {
	return &entities.Test{
		Id:   obj.Id,
		Name: obj.Name,
	}
}
