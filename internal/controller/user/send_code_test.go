package user

import (
	"bytes"
	"gin-study/internal/service/user"
	"gin-study/pkg/validate"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SendCodeSuite struct {
	suite.Suite
	ctrl     *gomock.Controller
	userServ *user.MockServer
	userCont *Controller
}

func (s *SendCodeSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	s.ctrl = gomock.NewController(s.T())

	var err error
	binding.Validator, err = validate.NewValidator()
	s.Require().NoError(err, nil)

	s.userServ = user.NewMockServer(s.ctrl)
	s.userCont = NewController(s.userServ)
}

func (s *SendCodeSuite) TearDownSuite() {
	s.ctrl.Finish()
}

func (s *SendCodeSuite) TestSendCode() {
	s.userServ.EXPECT().SendCode(gomock.Any(), gomock.Eq(&user.SendCodeRequest{
		Mobile: "13444444444",
		Type:   user.SendCodeTypeLogin,
	})).Return(&user.SendCodeResponse{Code: "123456"}, nil)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	body := bytes.NewBufferString(`{"mobile":"13444444444", "type":"login"}`)
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/v1/user/send/code", body)
	ctx.Params = gin.Params{{"mobile", "13444444444"}, {"code", "123456"}}
	ctx.Request.Header.Set("Content-Type", "application/json")

	s.userCont.SendCode(ctx)
	assert.Equal(s.T(), http.StatusOK, ctx.Writer.Status(), "状态返回应该为200")
}

func TestSendCodeSuite(t *testing.T) {
	suite.Run(t, new(SendCodeSuite))
}
