package user

import (
	"bytes"
	"gin-study/internal/service/user"
	"gin-study/pkg/validate"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var err error
	binding.Validator, err = validate.NewValidator()
	if err != nil {
		t.Fatal(err)
	}

	mockService := user.NewMockServer(ctrl)
	mockService.EXPECT().SendCode(gomock.Any(), gomock.Eq(&user.SendCodeRequest{
		Mobile: "13444444444",
		Type:   user.SendCodeTypeLogin,
	})).Return(&user.SendCodeResponse{Code: "123456"}, nil)

	userController := NewController(mockService)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	body := bytes.NewBufferString(`{"mobile":"13444444444", "type":"login"}`)
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/v1/user/send/code", body)
	ctx.Params = gin.Params{{"mobile", "13444444444"}, {"code", "123456"}}
	ctx.Request.Header.Set("Content-Type", "application/json")

	userController.SendCode(ctx)
	assert.Equal(t, http.StatusOK, ctx.Writer.Status(), "状态返回应该为200")

}
