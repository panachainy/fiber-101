package products

import (
	"fmt"
	"io"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// func setupTestSuite(tb testing.TB) func(tb testing.TB) {
// 	os.Setenv("REDIS_HOST", "127.0.0.1")
// 	os.Setenv("REDIS_PORT", "6379")
// 	os.Setenv("REDIS_PASSWORD", "redispassword1234")
// 	os.Setenv("REDIS_DB", "0")
// 	os.Setenv("MOOMALL_PLATFORM_URL", "http://127.0.0.1:3500")
// 	os.Setenv("MOOMALL_PLATFORM_NAME", "SALEPAGE_TEST")
// 	os.Setenv("MOOMALL_PLATFORM_SECRET", "TEST")

// 	config.Load("../example.env")

// 	database.InitRedis()

// 	return func(tb testing.TB) {
// 		logrus.Info("Teardown suite test")
// 	}
// }

// func NewMockProductRepo() ProductRepository {
// 	return &mockProductRepoImp{}
// }

// type mockProductRepoImp struct {
// 	DBConn *gorm.DB
// }

// func (r *mockProductRepoImp) GetAll() []Product {
// 	products := []Product{
// 		{
// 			Code: "xx",
// 		},
// 	}

// 	return products
// }

func TestGetGetsFunc(t *testing.T) {
	type args struct {
		p ProductRepository
	}
	tests := []struct {
		name string
		args
		expected string
	}{
		{
			name:     "when_get_Gets_func_should_return_func",
			expected: "{\"data\":[{\"ID\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"Code\":\"xx\",\"Price\":0,\"PriceDetailJa\":0}]}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := NewMockProductRepository(ctrl)
			mockRepo.EXPECT().GetAll().Return([]Product{
				{
					Code: "xx",
				},
			})

			if got := GetGetsFunc(mockRepo); !reflect.DeepEqual(got, tt.expected) {
				appMock := fiber.New()
				appMock.Get("/gets", got)

				req := httptest.NewRequest("GET", "/gets", nil)
				resp, err := appMock.Test(req, 5000)

				fmt.Println("=====================================")
				fmt.Println(err)

				fmt.Print(resp)
				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					assert.Equal(t, false, true, tt.name)
				}

				bodyString := string(bodyBytes)
				assert.Equal(t, tt.expected, bodyString)
			}
		})
	}
}
