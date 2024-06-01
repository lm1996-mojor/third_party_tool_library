package alibaba

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
)

// CreateClient
/**
 * API文档地址：https://next.api.aliyun.com/api-tools/sdk/Dysmsapi?version=2017-05-25&language=go-tea
 * 使用AK&SK初始化账号Client
 * 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例使用环境变量获取 AccessKey 的方式进行调用，仅供参考，
 * 建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
 * @param accessKeyId 访问密钥id
 * @param accessKeySecret 访问秘钥凭证
 * @return Client 访问客户端
 * @throws Exception 返回异常信息
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (client *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	client = &dysmsapi20170525.Client{}
	client, _err = dysmsapi20170525.NewClient(config)
	return client, _err
}
