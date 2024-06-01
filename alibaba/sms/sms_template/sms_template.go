package sms_template

import (
	"errors"

	"third_party_tool_library"
	"third_party_tool_library/alibaba"

	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
)

// AddSmsTemplate
/** 申请短信模板
 * @param accessKeyId 访问秘钥ID
 * @param accessKeySecret 访问秘钥凭证
 * @param templateName 模板名称，长度不超过 30 个字符。
 * @param templateContent 模板内容，长度不超过 500 个字符。更多规范，请参见: https://help.aliyun.com/document_detail/108253.html?spm=api-workbench.API%20Document.0.0.791d5513BOcVBx
 * @param remark 短信模板申请说明，是模板审核的参考信息之一。长度不超过 100 个字符。
 * @param templateType 短信类型
						0：验证码。
						1：短信通知。
						2：推广短信。
						3：国际/港澳台消息。
 * @return httpStatusCode 接口响应编码，包含参数检测的错误编码（可以判断该编码是否为200）
 * @return _resultMsg 响应对象（第三方返回的响应信息都在里面，包含业务错误）
 * @return error 错误响应对象（通常都是系统中的错误，业务错误不在其中，建议首先检测该对象是否 err == nil）
 * @return templateCode 短信模板 Code
*/
func AddSmsTemplate(accessKeyId, accessKeySecret, templateName, templateContent, remark string, templateType int32) (httpStatusCode int32, _resultMsg third_party_tool_library.ResponseResult, templateCode string, error error) {
	if templateName == "" {
		return 500, third_party_tool_library.ResponseResult{}, "", errors.New("短信模板名称不能为空")
	}
	if templateContent == "" {
		return 500, third_party_tool_library.ResponseResult{}, "", errors.New("短信模板内容不能为空")
	}
	if remark == "" || len(remark) > 100 {
		return 500, third_party_tool_library.ResponseResult{}, "", errors.New("短信模板申请说明不能为空，且不能超过100个字符")
	}
	if templateType < 0 || templateType > 3 {
		return 500, third_party_tool_library.ResponseResult{}, "", errors.New("短信类型不规范：0：验证码。\n1：短信通知。\n2：推广短信。\n3：国际/港澳台消息。")
	}
	client, _err := alibaba.CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return 500, third_party_tool_library.ResponseResult{}, "", _err
	}
	result, err := client.AddSmsTemplate(&dysmsapi20170525.AddSmsTemplateRequest{
		TemplateName:    &templateName,
		TemplateContent: &templateContent,
		TemplateType:    &templateType,
		Remark:          &remark,
	})
	if err != nil {
		return 500, third_party_tool_library.ResponseResult{}, "", _err
	}
	resp := third_party_tool_library.ResponseResult{Code: result.Body.Code, Message: result.Body.Message}
	return tea.Int32Value(result.StatusCode), resp, tea.StringValue(result.Body.TemplateCode), nil
}

// QuerySmsTemplateList
/** 查询短信模板列表
 * @param pageIndex 展示第几页的模板信息。默认取值为 1。
 * @param pageSize 每页展示的模板个数。默认取值为 10 最大 50。
 * @param accessKeyId 访问秘钥ID
 * @param accessKeySecret 访问秘钥凭证
 * @return httpStatusCode 接口响应编码，包含参数检测的错误编码（可以判断该编码是否为200）
 * @return _resultMsg 响应对象（第三方返回的响应信息都在里面，包含业务错误）
 * @return error 错误响应对象（通常都是系统中的错误，业务错误不在其中，建议首先检测该对象是否 err == nil）
 * @return smsTemplateList 短信模板列表。
 */
func QuerySmsTemplateList(accessKeyId, accessKeySecret string, pageIndex, pageSize int32) (httpStatusCode int32, _resultMsg third_party_tool_library.ResponseResult, smsTemplateList []*dysmsapi20170525.QuerySmsTemplateListResponseBodySmsTemplateList, error error) {
	if pageIndex <= 0 {
		pageIndex = 1
	}
	if pageSize < 10 {
		pageSize = 10
	}
	// 创建客户端对象
	client, _err := alibaba.CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return 500, third_party_tool_library.ResponseResult{}, nil, _err
	}
	result, err := client.QuerySmsTemplateList(&dysmsapi20170525.QuerySmsTemplateListRequest{PageIndex: &pageIndex, PageSize: &pageSize})
	if err != nil {
		return 500, third_party_tool_library.ResponseResult{}, nil, _err
	}
	resp := third_party_tool_library.ResponseResult{Code: result.Body.Code, Message: result.Body.Message}

	return tea.Int32Value(result.StatusCode), resp, result.Body.SmsTemplateList, nil
}

// QuerySmsTemplate
/** 查询短信模板的审核状态
 * @param templateCode 短信模板 Code
 * @param accessKeyId 访问秘钥ID
 * @param accessKeySecret 访问秘钥凭证
 * @return httpStatusCode 接口响应编码，包含参数检测的错误编码（可以判断该编码是否为200）
 * @return _resultMsg 响应对象（第三方返回的响应信息都在里面，包含业务错误）
 * @return error 错误响应对象（通常都是系统中的错误，业务错误不在其中，建议首先检测该对象是否 err == nil）
 * @return templateStatus 模板审核状态。取值：
										0：审核中。
										1：审核通过。
										2：审核未通过，请在返回参数 Reason 中查看审核失败原因。
										10：取消审核。
 * @return reason 审核备注。非验证码类型短信，请选择短信通知类型为推广短信。
						如果审核状态为审核通过或审核中，参数 Reason 显示为“无审批备注”。
						如果审核状态为审核未通过，参数 Reason 显示审核的具体原因。
*/
func QuerySmsTemplate(accessKeyId, accessKeySecret, templateCode string) (httpStatusCode int32, _resultMsg third_party_tool_library.ResponseResult, templateStatus int32, reason string, error error) {
	if templateCode == "" {
		return 500, third_party_tool_library.ResponseResult{}, 0, "", errors.New("短信模板编码不能为空")
	}
	// 创建客户端对象
	client, _err := alibaba.CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return 500, third_party_tool_library.ResponseResult{}, 0, "", _err
	}
	result, err := client.QuerySmsTemplate(&dysmsapi20170525.QuerySmsTemplateRequest{TemplateCode: &templateCode})
	if err != nil {
		return 500, third_party_tool_library.ResponseResult{}, 0, "", _err
	}
	resp := third_party_tool_library.ResponseResult{Code: result.Body.Code, Message: result.Body.Message}

	return tea.Int32Value(result.StatusCode), resp, tea.Int32Value(result.Body.TemplateStatus), tea.StringValue(result.Body.Reason), nil
}

// ModifySmsTemplate
/** 修改审核未通过的短信模板
 * @param templateCode 短信模板 Code
 * @param templateName 模板名称，长度不超过 30 个字符。
 * @param templateContent 模板内容，长度不超过 500 个字符。更多规范，请参见: https://help.aliyun.com/document_detail/108253.html?spm=api-workbench.API%20Document.0.0.791d5513BOcVBx
 * @param remark 短信模板申请说明，是模板审核的参考信息之一。长度不超过 100 个字符。
 * @param templateType 短信类型
						0：验证码。
						1：短信通知。
						2：推广短信。
						3：国际/港澳台消息。
 * @param accessKeyId 访问秘钥ID
 * @param accessKeySecret 访问秘钥凭证
 * @return httpStatusCode 接口响应编码，包含参数检测的错误编码（可以判断该编码是否为200）
 * @return _resultMsg 响应对象（第三方返回的响应信息都在里面，包含业务错误）
 * @return error 错误响应对象（通常都是系统中的错误，业务错误不在其中，建议首先检测该对象是否 err == nil）
 * @return templateCode 短信模板 Code
*/
func ModifySmsTemplate(accessKeyId, accessKeySecret, templateCode, templateName, templateContent, remark string, templateType int32) (httpStatusCode int32, _resultMsg third_party_tool_library.ResponseResult, respTemplateCode string, error error) {
	if templateType < 0 || templateType > 3 {
		return 500, third_party_tool_library.ResponseResult{}, "", errors.New("短信类型不规范：0：验证码。\n1：短信通知。\n2：推广短信。\n3：国际/港澳台消息。")
	}
	if templateCode == "" {
		return 500, third_party_tool_library.ResponseResult{}, "", errors.New("短信模板编码不能为空")
	}
	// 创建客户端对象
	client, _err := alibaba.CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return 500, third_party_tool_library.ResponseResult{}, "", _err
	}
	// 创建并初始化发送信息对象

	result, err := client.ModifySmsTemplate(&dysmsapi20170525.ModifySmsTemplateRequest{
		TemplateCode:    &templateCode,
		Remark:          &remark,
		TemplateType:    &templateType,
		TemplateContent: &templateContent,
		TemplateName:    &templateName,
	})
	if err != nil {
		return 500, third_party_tool_library.ResponseResult{}, "", _err
	}
	resp := third_party_tool_library.ResponseResult{Code: result.Body.Code, Message: result.Body.Message}
	return tea.Int32Value(result.StatusCode), resp, tea.StringValue(result.Body.TemplateCode), nil
}

// DeleteSmsTemplate
/** 删除短信模板
 * @param templateCode 短信模板 Code
 * @param accessKeyId 访问秘钥ID
 * @param accessKeySecret 访问秘钥凭证
 * @return httpStatusCode 接口响应编码，包含参数检测的错误编码（可以判断该编码是否为200）
 * @return _resultMsg 响应对象（第三方返回的响应信息都在里面，包含业务错误）
 * @return error 错误响应对象（通常都是系统中的错误，业务错误不在其中，建议首先检测该对象是否 err == nil）
 * @param templateCode 短信模板 Code
 */
func DeleteSmsTemplate(accessKeyId, accessKeySecret, templateCode string) (httpStatusCode int32, _resultMsg third_party_tool_library.ResponseResult, respTemplateCode string, error error) {
	if templateCode == "" {
		return 500, third_party_tool_library.ResponseResult{}, "", errors.New("短信模板编码不能为空")
	}
	// 创建客户端对象
	client, _err := alibaba.CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return 500, third_party_tool_library.ResponseResult{}, "", _err
	}
	// 创建并初始化发送信息对象

	result, err := client.DeleteSmsTemplate(&dysmsapi20170525.DeleteSmsTemplateRequest{
		TemplateCode: &templateCode,
	})
	if err != nil {
		return 500, third_party_tool_library.ResponseResult{}, "", _err
	}
	resp := third_party_tool_library.ResponseResult{Code: result.Body.Code, Message: result.Body.Message}
	return tea.Int32Value(result.StatusCode), resp, tea.StringValue(result.Body.TemplateCode), nil
}
