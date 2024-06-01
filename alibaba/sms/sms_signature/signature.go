package sms_signature

import (
	"errors"
	"strings"

	"lm1996-mojor/third_party_tool_library"
	"lm1996-mojor/third_party_tool_library/alibaba"

	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
)

// AddSmsSignature
/** 申请短信签名
 * @param accessKeyId 访问秘钥ID
 * @param accessKeySecret 访问秘钥凭证
 * @param signName 签名名称
					短信签名申请说明，长度不超过 200 个字符。 场景说明是签名审核的参考信息之一。请详细描述已上线业务的使用场景，并提供可以验证这些业务的网站链接、
					已备案域名地址、应用市场下载链接、公众号或小程序全称等信息。对于登录场景，还需提供测试账号密码。信息完善的申请说明会提高签名、模板的审核效率。
 * @param remark 申请理由
 * @param signSource 签名来源（0：企事业单位的全称或简称。1：工信部 备案网站的全称或简称。2：App 应用的全称或简称。3：公众号或小程序的全称或简称。4：电商平台店铺名的全称或简称。5：商标名的全称或简称。）
 * @param signType 签名类型 （0 验证码 1 通用）
 * @param signFileList 签名文件列表。如果签名用途为他用或个人认证用户的自用签名来源为企事业单位名时，还需上传证明文件和委托授权书，详情请参见证明文件和授权委托书。
	- @param signFileList[FileContents] 签名的资质证明文件经 base64 编码后的字符串。图片不超过 2 MB
	- @param signFileList[FileSuffix] 签名的证明文件格式，支持上传多张图片。当前支持 JPG、PNG、GIF 或 JPEG 格式的图片
 * @return httpStatusCode 接口响应编码，包含参数检测的错误编码（可以判断该编码是否为200）
 * @return _resultMsg 响应对象（第三方返回的响应信息都在里面，包含业务错误）
 * @return error 错误响应对象（通常都是系统中的错误，业务错误不在其中，建议首先检测该对象是否 err == nil）
*/
func AddSmsSignature(accessKeyId, accessKeySecret, signName, remark string, signSource, signType int32, signFileList []*dysmsapi20170525.AddSmsSignRequestSignFileList) (httpStatusCode int32, _resultMsg third_party_tool_library.ResponseResult, error error) {
	if signSource < 0 || signSource > 5 {
		return 400, third_party_tool_library.ResponseResult{}, errors.New("短信签名来源不规范：0：企事业单位的全称或简称。\n1：工信部备案网站的全称或简称。\n2：App 应用的全称或简称。\n3：公众号或小程序的全称或简称。\n4：电商平台店铺名的全称或简称。\n5：商标名的全称或简称。")
	}
	if signType < 1 || signType > 2 {
		return 400, third_party_tool_library.ResponseResult{}, errors.New("短信签名类型不规范：0：验证码\n1：通用")
	}
	if signName != "" {
		signName = strings.Trim(signName, " ")
		if len(signName) > 12 {
			return 400, third_party_tool_library.ResponseResult{}, errors.New("短信签名名称长度不能超过12个字符")
		}
	} else {
		return 400, third_party_tool_library.ResponseResult{}, errors.New("短信签名名称不能为空")
	}

	remark = strings.Trim(remark, " ")
	if len(remark) > 200 {
		return 400, third_party_tool_library.ResponseResult{}, errors.New("短信签名申请说明，长度不能超过 200 个字符")
	}
	if len(signFileList) < 0 {
		return 400, third_party_tool_library.ResponseResult{}, errors.New("短信签名申请证明文件不能为空")
	}
	// 创建客户端对象
	client, _err := alibaba.CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return 500, third_party_tool_library.ResponseResult{}, _err
	}
	// 创建并初始化发送信息对象

	result, err := client.AddSmsSign(&dysmsapi20170525.AddSmsSignRequest{
		SignName:     &signName,
		Remark:       &remark,
		SignSource:   &signSource,
		SignType:     &signType,
		SignFileList: signFileList,
	})
	if err != nil {
		return 500, third_party_tool_library.ResponseResult{}, _err
	}
	resp := third_party_tool_library.ResponseResult{Code: result.Body.Code, Message: result.Body.Message}
	return tea.Int32Value(result.StatusCode), resp, nil
}

// QuerySmsSignList
/** 查询短信签名列表
 * @param accessKeyId
 * @param accessKeySecret
 * @param pageIndex
 * @param pageSize
 * @return httpStatusCode 接口响应编码，包含参数检测的错误编码（可以判断该编码是否为200）
 * @return _resultMsg 响应对象（第三方返回的响应信息都在里面，包含业务错误）
 * @return smsSignList 短信签名列表,对象详细情况：https://next.api.aliyun.com/document/Dysmsapi/2017-05-25/QuerySmsSignList?accounttraceid=276630863bf5478da6f466bbf658d919yrps
 * @return error 错误响应对象（通常都是系统中的错误，业务错误不在其中，建议首先检测该对象是否 err == nil）
 */
func QuerySmsSignList(accessKeyId, accessKeySecret string, pageIndex, pageSize int32) (httpStatusCode int32, _resultMsg third_party_tool_library.ResponseResult, smsSignList []*dysmsapi20170525.QuerySmsSignListResponseBodySmsSignList, error error) {
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
	result, err := client.QuerySmsSignList(&dysmsapi20170525.QuerySmsSignListRequest{PageIndex: &pageIndex, PageSize: &pageSize})
	if err != nil {
		return 500, third_party_tool_library.ResponseResult{}, nil, _err
	}
	resp := third_party_tool_library.ResponseResult{Code: result.Body.Code, Message: result.Body.Message}

	return tea.Int32Value(result.StatusCode), resp, result.Body.SmsSignList, nil
}

// ModifySmsSign
/** 修改短信签名
 * @param accessKeyId 访问秘钥ID
 * @param accessKeySecret 访问秘钥凭证
 * @param signName 签名名称
					短信签名申请说明，长度不超过 200 个字符。 场景说明是签名审核的参考信息之一。请详细描述已上线业务的使用场景，并提供可以验证这些业务的网站链接、
					已备案域名地址、应用市场下载链接、公众号或小程序全称等信息。对于登录场景，还需提供测试账号密码。信息完善的申请说明会提高签名、模板的审核效率。
 * @param remark 申请理由
 * @param signSource 签名来源（0：企事业单位的全称或简称。1：工信部 备案网站的全称或简称。2：App 应用的全称或简称。3：公众号或小程序的全称或简称。4：电商平台店铺名的全称或简称。5：商标名的全称或简称。）
 * @param signType 签名类型 （0 验证码 1 通用）
 * @param signFileList [dysmsapi20170525.ModifySmsSignRequest] 签名文件列表。如果签名用途为他用或个人认证用户的自用签名来源为企事业单位名时，还需上传证明文件和委托授权书，详情请参见证明文件和授权委托书。
	- @param signFileList[FileContents] 签名的资质证明文件经 base64 编码后的字符串。图片不超过 2 MB
	- @param signFileList[FileSuffix] 签名的证明文件格式，支持上传多张图片。当前支持 JPG、PNG、GIF 或 JPEG 格式的图片
 * @return httpStatusCode 接口响应编码，包含参数检测的错误编码（可以判断该编码是否为200）
 * @return _resultMsg 响应对象（第三方返回的响应信息都在里面，包含业务错误）
 * @return error 错误响应对象（通常都是系统中的错误，业务错误不在其中，建议首先检测该对象是否 err == nil）
*/
func ModifySmsSign(accessKeyId, accessKeySecret, signName, remark string, signSource, signType int32, signFileList []*dysmsapi20170525.ModifySmsSignRequestSignFileList) (httpStatusCode int32, _resultMsg third_party_tool_library.ResponseResult, error error) {
	if signSource < 0 || signSource > 5 {
		return 400, third_party_tool_library.ResponseResult{}, errors.New("短信签名来源不规范：0：企事业单位的全称或简称。\n1：工信部备案网站的全称或简称。\n2：App 应用的全称或简称。\n3：公众号或小程序的全称或简称。\n4：电商平台店铺名的全称或简称。\n5：商标名的全称或简称。")
	}
	if signType < 1 || signType > 2 {
		return 400, third_party_tool_library.ResponseResult{}, errors.New("短信签名类型不规范：0：验证码\n1：通用")
	}
	// 创建客户端对象
	client, _err := alibaba.CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return 500, third_party_tool_library.ResponseResult{}, _err
	}
	// 创建并初始化发送信息对象

	result, err := client.ModifySmsSign(&dysmsapi20170525.ModifySmsSignRequest{
		SignName:     &signName,
		Remark:       &remark,
		SignSource:   &signSource,
		SignType:     &signType,
		SignFileList: signFileList,
	})
	if err != nil {
		return 500, third_party_tool_library.ResponseResult{}, _err
	}
	resp := third_party_tool_library.ResponseResult{Code: result.Body.Code, Message: result.Body.Message}
	return tea.Int32Value(result.StatusCode), resp, nil
}

// DeleteSmsSign
/** 删除短信签名
 * @param accessKeyId 访问秘钥ID
 * @param accessKeySecret 访问秘钥凭证
 * @param signName 签名名称
 * @return httpStatusCode 接口响应编码，包含参数检测的错误编码（可以判断该编码是否为200）
 * @return _resultMsg 响应对象（第三方返回的响应信息都在里面，包含业务错误）
 * @return error 错误响应对象（通常都是系统中的错误，业务错误不在其中，建议首先检测该对象是否 err == nil）
 */
func DeleteSmsSign(accessKeyId, accessKeySecret, signName string) (httpStatusCode int32, _resultMsg third_party_tool_library.ResponseResult, error error) {
	if signName == "" {
		return 400, third_party_tool_library.ResponseResult{}, errors.New("短信签名名称不能为空")
	}
	// 创建客户端对象
	client, _err := alibaba.CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return 500, third_party_tool_library.ResponseResult{}, _err
	}
	result, err := client.DeleteSmsSign(&dysmsapi20170525.DeleteSmsSignRequest{
		SignName: &signName,
	})
	if err != nil {
		return 500, third_party_tool_library.ResponseResult{}, _err
	}
	resp := third_party_tool_library.ResponseResult{Code: result.Body.Code, Message: result.Body.Message}
	return tea.Int32Value(result.StatusCode), resp, nil
}

// QuerySmsSign
/** 查询短信签名申请状态
 * @param accessKeyId 访问秘钥ID
 * @param accessKeySecret 访问秘钥凭证
 * @param signName 签名名称
 * @return httpStatusCode 接口响应编码，包含参数检测的错误编码（可以判断该编码是否为200）
 * @return _resultMsg 响应对象（第三方返回的响应信息都在里面，包含业务错误）
 * @return auditStatus 签名审核状态
					0：审核中。1：审核通过。2：审核失败，请在返回参数 Reason 中查看审核失败原因。10：取消审核。
 * @return reason 审核备注。
					如果审核状态为审核通过或审核中，参数 Reason 显示为“无审核备注”。
					如果审核状态为审核未通过，参数 Reason 显示审核的具体原因。
 * @return error 错误响应对象（通常都是系统中的错误，业务错误不在其中，建议首先检测该对象是否 err == nil）
*/
func QuerySmsSign(accessKeyId, accessKeySecret, signName string) (httpStatusCode int32, _resultMsg third_party_tool_library.ResponseResult, auditStatus int32, reason string, error error) {
	if signName == "" {
		return 400, third_party_tool_library.ResponseResult{}, -1, "", errors.New("短信签名名称不能为空")
	}
	// 创建客户端对象
	client, _err := alibaba.CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return 500, third_party_tool_library.ResponseResult{}, -1, "", _err
	}
	result, err := client.QuerySmsSign(&dysmsapi20170525.QuerySmsSignRequest{
		SignName: &signName,
	})
	if err != nil {
		return 500, third_party_tool_library.ResponseResult{}, -1, "", _err
	}
	resp := third_party_tool_library.ResponseResult{Code: result.Body.Code, Message: result.Body.Message}
	return tea.Int32Value(result.StatusCode), resp, tea.Int32Value(result.Body.SignStatus), tea.StringValue(result.Body.Reason), nil
}
