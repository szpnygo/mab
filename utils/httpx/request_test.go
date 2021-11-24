package httpx

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type ErrorResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type RequestData struct {
	AppID string `json:"appid"`
}

const GET_REQUEST_URL = "https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo"

const POST_REQUEST_URL = "https://api.weixin.qq.com/cgi-bin/clear_quota"

func Test_request_Request(t *testing.T) {
	Convey("init", t, func() {

		Convey("get", func() {
			response, err := Get(GET_REQUEST_URL).Debug().Params(map[string]string{
				"access_token": "access_token",
			}).Request()

			So(err, ShouldBeNil)
			So(response, ShouldNotBeNil)

			var errorResponse ErrorResponse
			err = response.JSON(&errorResponse)

			So(err, ShouldBeNil)
			So(errorResponse.ErrCode, ShouldEqual, 40001)

		})

		Convey("get add params", func() {
			response, err := Get(GET_REQUEST_URL).Debug().AddParam("access_token", "token").Request()

			So(err, ShouldBeNil)
			So(response, ShouldNotBeNil)

			var errorResponse ErrorResponse
			err = response.JSON(&errorResponse)

			So(err, ShouldBeNil)
			So(errorResponse.ErrCode, ShouldEqual, 40001)

		})

		Convey("post", func() {
			requestData := RequestData{
				AppID: "wx448f04719cd48f69",
			}

			var errorResponse2 ErrorResponse

			err := Post(GET_REQUEST_URL).Debug().Params(map[string]string{
				"access_token": "access_token",
			}).Struct(&requestData).JSON(&errorResponse2)

			So(err, ShouldBeNil)
			So(errorResponse2, ShouldNotBeNil)
			So(errorResponse2.ErrCode, ShouldEqual, 40001)

		})

	})
}
