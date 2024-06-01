package sms_execute

import (
	"lm1996-mojor/third_party_tool_library"
	"lm1996-mojor/third_party_tool_library/alibaba"

	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// SmsSend 短信发送
/**
 * 批量发送使用的是同一个短信模板，手机号码、短信签名、模板中的参数、都是json串的方式，都是一一对应的，例如：phoneNumbers{"139xxx1","136xxx1"},signName{"xxx通知","xxx短信"}
 * 单个短信发送是每个号码都将使用不同的短信模板和签名
 * 错误码列表: https://help.aliyun.com/zh/sms/developer-reference/api-error-codes
 * @param phoneNumbers 接收对象的手机号码
 * @param signName 短信签名名称
 * @param templateCode  短信模板编号
 * @param templateParam 短信模板中的参数
 * @param isBatchSend 是否进行批量发送
 * @return int32 接口响应编码，包含参数检测的错误编码（可以判断该编码是否为200）
 * @return third_party_tool_library.ResponseResult 响应对象（第三方返回的响应信息都在里面，包含业务错误）
 * @return error 错误响应对象（通常都是系统中的错误，业务错误不在其中，建议首先检测该对象是否 err == nil）
 */
func SmsSend(accessKeyId, accessKeySecret, phoneNumbers, signName, templateCode, templateParam string, isBatchSend bool) (int32, third_party_tool_library.ResponseResult, error) {
	var statusCode int32
	// 创建客户端对象
	client, _err := alibaba.CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return statusCode, third_party_tool_library.ResponseResult{}, _err
	}
	// 创建并初始化发送信息对象
	respMsg := third_party_tool_library.ResponseResult{}
	if isBatchSend {
		sendBatchSmsRequest := &dysmsapi20170525.SendBatchSmsRequest{
			PhoneNumberJson:   tea.String(phoneNumbers),
			SignNameJson:      tea.String(signName),
			TemplateCode:      tea.String(templateCode),
			TemplateParamJson: tea.String(templateParam),
		}
		statusCode, respMsg, _err = batchSmsSend(client, sendBatchSmsRequest, &util.RuntimeOptions{})
	} else {
		sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
			PhoneNumbers:  tea.String(phoneNumbers),
			SignName:      tea.String(signName),
			TemplateCode:  tea.String(templateCode),
			TemplateParam: tea.String(templateParam),
		}
		statusCode, respMsg, _err = singleSmsSend(client, sendSmsRequest, &util.RuntimeOptions{})
	}
	if _err != nil {
		return statusCode, third_party_tool_library.ResponseResult{}, _err
	}
	return statusCode, respMsg, nil
}

// 发送单个短信
func singleSmsSend(client *dysmsapi20170525.Client, req *dysmsapi20170525.SendSmsRequest, runtime *util.RuntimeOptions) (int32, third_party_tool_library.ResponseResult, error) {
	result, err := client.SendSmsWithOptions(req, runtime)
	if err != nil {
		return 500, third_party_tool_library.ResponseResult{}, err
	}
	return tea.Int32Value(result.StatusCode), third_party_tool_library.NewResult(result.Body.Code, result.Body.Message), nil
}

// 批量发送短信
func batchSmsSend(client *dysmsapi20170525.Client, req *dysmsapi20170525.SendBatchSmsRequest, runtime *util.RuntimeOptions) (int32, third_party_tool_library.ResponseResult, error) {
	result, err := client.SendBatchSmsWithOptions(req, runtime)
	if err != nil {
		return 500, third_party_tool_library.ResponseResult{}, err
	}
	return tea.Int32Value(result.StatusCode), third_party_tool_library.NewResult(result.Body.Code, result.Body.Message), nil
}
