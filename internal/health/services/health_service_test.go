package services_test

import (
	"testing"

	health_svc "github.com/HenCor2019/book-file-api/internal/health/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDependencyService struct {
	mock.Mock
}

func (m *MockDependencyService) SomeMethod() string {
	args := m.Called()
	return args.String(0)
}

func TestCheckHealth(t *testing.T) {
	service := health_svc.New()

	t.Run("Should return OK", func(t *testing.T) {
		t.Parallel()
		result := service.CheckHealth()
		assert.Equal(t, "ok", result)
	})
}

func TestHelloWorld(t *testing.T) {
	service := health_svc.New()

	t.Run("Should return Hello World from ENV: ", func(t *testing.T) {
		t.Parallel()
		result := service.HelloWorld()
		assert.Equal(t, "Hello World from ENV: ", result)
	})
}
