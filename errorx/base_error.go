package errorx

import (
	"fmt"
)

type CodeErrorResponse struct {
	TraceID   string                   `json:"TraceID"`
	ErrorData CodeErrorResponseContent `json:"Error"`
	Data      interface{}              `json:"Data,omitempty"`
	Token     interface{}              `json:"Token"`
}

type CodeErrorResponseContent struct {
	Code ErrCode `json:"Code"`
	Msg  string  `json:"Msg"`
}

type ResponseListData struct {
	List []interface{} `json:"List"` // 确保 List 是一个数组
}

func (o ErrCode) GenResonse(traceID string, msg string) CodeErrorResponse {
	targetMsg := msg
	if len(msg) < 1 {
		targetMsg = codeMsgMap[o]
	}
	return CodeErrorResponse{
		TraceID: traceID,
		ErrorData: CodeErrorResponseContent{
			Code: o,
			Msg:  targetMsg,
		},
	}
}

func (o ErrCode) GenResonseWithData(traceID string, data interface{}, token interface{}) CodeErrorResponse {
	return CodeErrorResponse{
		TraceID: traceID,
		ErrorData: CodeErrorResponseContent{
			Code: o,
			Msg:  codeMsgMap[o],
		},
		Data:  data,
		Token: token,
	}
}

func (o ErrCode) GenResponseWithEmptyListData(traceID string, token interface{}) CodeErrorResponse {
	return CodeErrorResponse{
		TraceID: traceID,
		ErrorData: CodeErrorResponseContent{
			Code: o,
			Msg:  codeMsgMap[o],
		},
		Data: ResponseListData{
			List: []interface{}{}, // 返回一个空数组
		},
		Token: token,
	}
}

func (o ErrCode) GetMsg() string {
	return codeMsgMap[o]
}

func (o ErrCode) GetCode() string {
	return o.GetCode()
}

// GenError
func (o ErrCode) GenError(msg ...string) *CodeErrorResponseContent {
	resp := CodeErrorResponseContent{
		Code: o,
		Msg:  codeMsgMap[o],
	}
	if len(msg) > 0 {
		resp.Msg = msg[0]
	}

	return &resp
}

// Error 错误打印
func (o *CodeErrorResponseContent) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", string(o.Code), o.Msg)
}

func (o *CodeErrorResponseContent) GenResonse(traceID string) CodeErrorResponse {
	return CodeErrorResponse{
		TraceID:   traceID,
		ErrorData: *o,
	}
}

type ErrCode string

const (
	ErrCodeNone                                    ErrCode = "Success"
	ErrCodeDataNotFound                            ErrCode = "DataNotFound"                                   // 数据不存在
	ErrCodeSkinDataNotFound                        ErrCode = "SkinDataNotFound"                               // 测肤数据不存在
	ErrCodeInternal                                ErrCode = "InternalError"                                  // 内部服务错误
	ErrCodeInvalidParamter                         ErrCode = "InvalidParamter"                                // 参数错误
	ErrCodeAuthorizationTokenValid                 ErrCode = "AuthorizationTokenValid"                        // token失效
	ErrCodeUserFrozen                              ErrCode = "UserForzen"                                     // 用户冻结
	ErrCodeVirifyCodeInvalid                       ErrCode = "InvalidVerifyCode"                              // 验证码已失效
	ErrCodeVerifyCodeErr                           ErrCode = "ErrCodeVirifyCodeErr"                           // 验证码错误
	ErrCodeDuplicateKey                            ErrCode = "DuplicateKey"                                   // 重复主键
	ErrPermissionDenied                            ErrCode = "PermissionDenied"                               //
	ErrCodeInEffectInternal                        ErrCode = "InEffect"                                       // 内部服务错误
	ErrCodePwdErr                                  ErrCode = "ErrCodePwdErr"                                  // 密码错误
	ErrCodeUserDataNotFound                        ErrCode = "UserDataNotFound"                               // 用户不存在
	ErrCodeIllegalUser                             ErrCode = "IllegalUser"                                    // 非法用户
	ErrCodeBoundUser                               ErrCode = "BoundUser"                                      // 已绑定用户，解绑后操作
	ErrCodeSourcePwdError                          ErrCode = "SourcePwdError"                                 // 修改密码，原始密码错误
	ErrCodeTooManyOperatins                        ErrCode = "TooManyOperations"                              // 操作次数过多
	ErrCodeTransgression                           ErrCode = "ErrCodeTransgression"                           // 越界操作
	ErrCodeArrayOverreach                          ErrCode = "ErrCodeArrayOverreach"                          // 数组越界
	ErrCodeValid                                   ErrCode = "ErrCodeValid"                                   // code失效
	ErrCodeMeiCeCalcErr                            ErrCode = "MeiCeCalcErr"                                   // 美测计算失败
	ErrCodeEmployeeStoreDataNotFound               ErrCode = "ErrCodeEmployeeStoreDataNotFound"               // 员工门店关系数据不存在
	ErrCodeEmployeeFrozen                          ErrCode = "ErrCodeEmployeeFrozen"                          // 员工已冻结
	ErrCodeShopFrozen                              ErrCode = "ErrCodeShopFrozen"                              // 门店已冻结
	ErrCodeEmployeeAndShopFrozen                   ErrCode = "ErrCodeEmployeeAndShopFrozen"                   // 账号&门店都已冻结
	ErrCodeWxOfficialAccountConfDataNotFound       ErrCode = "ErrCodeWxOfficialAccountConfDataNotFound"       // 微信公众号配置数据不存在
	ErrCodeWxOfficialAccountSubscribeErr           ErrCode = "ErrCodeWxOfficialAccountSubscribeErr"           // 微信公众号订阅错误
	ErrCodeWxOfficialAccountSearchSubscribe        ErrCode = "ErrCodeWxOfficialAccountSearchSubscribe"        // 自己搜索微信公众号手动关注
	ErrCodeGainWxComponentVerifyTicketErr          ErrCode = "ErrCodeGainWxComponentVerifyTicketErr"          // 获取微信WxComponentVerifyTicketErr
	ErrCodeGainWxComponentAccessTokenErr           ErrCode = "ErrCodeGainWxComponentAccessTokenErr"           // 获取微信WxComponentAccessTokenErr
	ErrCodeWxOfficialAccountSubscribeTemplateEmpty ErrCode = "ErrCodeWxOfficialAccountSubscribeTemplateEmpty" // 微信订阅模版为空
	ErrCodeThirdPartyInterfaceFailed               ErrCode = "ErrCodeThirdPartyInterfaceFailed"               // 第三方接口出错，请稍后再试
	ErrCodeOrderNotFound                           ErrCode = "ErrCodeOrderNotFound"                           // 正向订单不存在
	ErrCodeRefundOrderNotFound                     ErrCode = "ErrCodeRefundOrderNotFound"                     // 退款订单不存在
)

var (
	codeMsgMap = map[ErrCode]string{
		ErrCodeDataNotFound:                            "数据不存在",
		ErrCodeSkinDataNotFound:                        "测肤数据不存在",
		ErrCodeInternal:                                "数据库或服务异常，请刷新重试",
		ErrCodeInvalidParamter:                         "参数错误",
		ErrCodeAuthorizationTokenValid:                 "token失效",
		ErrCodeUserFrozen:                              "无法登录，该账号已被冻结",
		ErrCodeVirifyCodeInvalid:                       "验证码已失效",
		ErrCodeDuplicateKey:                            "数据已存在",
		ErrPermissionDenied:                            "获取权限失败，请联系管理员",
		ErrCodeVerifyCodeErr:                           "验证码错误",
		ErrCodeInEffectInternal:                        "验证码仍在生效中，请稍后重试",
		ErrCodePwdErr:                                  "密码错误",
		ErrCodeUserDataNotFound:                        "用户不存在",
		ErrCodeIllegalUser:                             "非法用户，请退出后重新登录，若有疑问请联系管理员或售后！",
		ErrCodeBoundUser:                               "已绑定用户，解绑后操作",
		ErrCodeSourcePwdError:                          "原始密码错误",
		ErrCodeTooManyOperatins:                        "操作次数超过限制",
		ErrCodeTransgression:                           "越界操作，请退出后重新登录，若有疑问请联系管理员或售后！",
		ErrCodeValid:                                   "Code失效",
		ErrCodeMeiCeCalcErr:                            "美测计算失败",
		ErrCodeEmployeeStoreDataNotFound:               "该员工信息或商家、门店信息不存在，无法登录",
		ErrCodeEmployeeFrozen:                          "无法登录，该员工账号已被冻结",
		ErrCodeShopFrozen:                              "无法登录，该门店已被冻结",
		ErrCodeEmployeeAndShopFrozen:                   "无法登录，该员工账号以及门店均已被冻结",
		ErrCodeArrayOverreach:                          "字段长度或数据类型超出限制",
		ErrCodeWxOfficialAccountConfDataNotFound:       "绑定微信粉丝关系失败，微信公众号配置数据不存在",
		ErrCodeWxOfficialAccountSubscribeErr:           "微信订阅失败，请取关公众号后-重新扫描关注",
		ErrCodeWxOfficialAccountSearchSubscribe:        "自己搜索微信公众号手动关注无效",
		ErrCodeGainWxComponentVerifyTicketErr:          "获取微信WxComponentVerifyTicket失败",
		ErrCodeGainWxComponentAccessTokenErr:           "获取微信WxComponentAccessToken失败",
		ErrCodeWxOfficialAccountSubscribeTemplateEmpty: "无微信消息模板，请在后台添加后重试",
		ErrCodeThirdPartyInterfaceFailed:               "第三方接口出错，请稍后再试",
		ErrCodeOrderNotFound:                           "主订单数据不存在",
		ErrCodeRefundOrderNotFound:                     "退款单数据不存在",
	}
)

var (
	CodeMsgType = map[interface{}]string{
		"trade_TradeClose":          "交易关闭",
		"trade_TradeCreate":         "交易创建",
		"trade_TradeBuyerPay":       "买家付款(即商家待发货)",
		"trade_TradePaid":           "交易支付",
		"trade_TradeSuccess":        "交易成功",
		"trade_refund_BuyerCreated": "买家发起退款",
	}
)
