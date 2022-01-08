package products

import (
	"fmt"
	"io"
	"net/http/httptest"
	reflect "reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetGetsFunc(t *testing.T) {
	type args struct {
		p func(ctrl *gomock.Controller) *MockProductRepository
	}
	tests := []struct {
		name string
		args
		expected string
	}{
		{
			name: "when_get_Gets_func_should_return_func",
			args: args{
				p: func(ctrl *gomock.Controller) *MockProductRepository {
					mockRepo := NewMockProductRepository(ctrl)
					mockRepo.EXPECT().GetAll().Return([]Product{
						{
							Code: "xx",
						},
					})

					return mockRepo
				},
			},
			expected: "{\"data\":[{\"ID\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"Code\":\"xx\",\"Price\":0,\"PriceDetailJa\":0}]}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := tt.args.p(ctrl)

			if got := GetGetsFunc(mockRepo); !reflect.DeepEqual(got, tt.expected) {
				appMock := fiber.New()
				appMock.Get("/gets", got)

				req := httptest.NewRequest("GET", "/gets", nil)
				resp, _ := appMock.Test(req, 5000)

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
