package main

import (
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	kms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/kms/v20190118"
)

func main() {
	credential := common.NewCredential(
		"IKIDqeKlxR8brszxFNSM9l9G7h15xc1C9uRk",
		"82jKisoY3RfoKp9JGxd5HExRf2Pckrre",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "kms.tencentcloudapi.com"
	client, _ := kms.NewClient(credential, "ap-hongkong", cpf)

	request := kms.NewDecryptRequest()

	request.CiphertextBlob = common.StringPtr("BYCUP6xePY1vS9u2xF2LoRGV4P7r7G5BwQ3rfaPvZaTZ7yU7IsgU8152hH8iPj8mn57MtKQrcX6fTNN9B11qPg==-k-fKVP3WIlGpg8m9LMW4jEkQ==-k-GkK58mChGcT4Re4N6NVDnh+VSlX7/9K396K27ZgaoUhF7VbO25xhA0h6VaoGSqseJeSzBX/yTapgpx40B9yzLG0vADlm+PCUnNbV6737NvmB16OefVxs9p/RB74F1NUL+60UPA==")

	response, err := client.Decrypt(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
