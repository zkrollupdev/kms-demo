package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	kms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/kms/v20190118"
)

func main() {

	generateCMK()
	generate_data_key()
	encrypt_data()
	dencrypt_data()

	//base64 encode example
	input := "fc373de160b2028ff48d7df7eb16b8a4fcf1761ce4734ae42fd21e8f34f9b97f"
	encodeString := base64.StdEncoding.EncodeToString([]byte(input))
	fmt.Println(encodeString)

	//base64 decode example
	decodeString, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(decodeString))

}

func generateCMK() {
	credential := common.NewCredential( //your api access key
		"IKIDqeKlxR8brszxFNSM9l9G7h15xc1C9uRk",
		"82jKisoY3RfoKp9JGxd5HExRf2Pckrre",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "kms.tencentcloudapi.com"
	client, _ := kms.NewClient(credential, "ap-hongkong", cpf)
	request_1 := kms.NewCreateKeyRequest()
	request_1.Alias = common.StringPtr("test2")
	response, err := client.CreateKey(request_1)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s\n", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("generateCMK:%s\n", response.ToJsonString())
}

//{"Response":{"KeyId":"36a21228-ea3a-11ed-9810-5269be2b21c4","Alias":"test3","CreateTime":1683177126,"Description":"","KeyState":"Enabled","KeyUsage":"ENCRYPT_DECRYPT","TagCode":0,"TagMsg":"","HsmClusterId":"","RequestId":"5802daf6-bd0d-4c91-afee-8fbdd44f2ec7"}}

func generate_data_key() {
	credential := common.NewCredential( //your api access key
		"IKIDqeKlxR8brszxFNSM9l9G7h15xc1C9uRk",
		"82jKisoY3RfoKp9JGxd5HExRf2Pckrre",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "kms.tencentcloudapi.com"
	client, _ := kms.NewClient(credential, "ap-hongkong", cpf)

	request := kms.NewGenerateDataKeyRequest()

	request.KeyId = common.StringPtr("1fd7f96c-ea37-11ed-b68d-32c48ca45083") //your kms key_id,replace it
	request.KeySpec = common.StringPtr("AES_256")

	response, err := client.GenerateDataKey(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("generate_data_key:%s\n", response.ToJsonString())
}

// {"Response":{"KeyId":"1fd7f96c-ea37-11ed-b68d-32c48ca45083","Plaintext":"yFHQWVkcOCjLHvdXG/vG3zcwROJHqj0x23SaC4ncR/c=","CiphertextBlob":"BYCUP6xePY1vS9u2xF2LoRGV4P7r7G5BwQ3rfaPvZaTZ7yU7IsgU8152hH8iPj8mn57MtKQrcX6fTNN9B11qPg==-k-fKVP3WIlGpg8m9LMW4jEkQ==-k-J+NJr09sk7kmOhDV+TJSF+9vOyypw8n5xzLeut2NFo5yTsAuhqNBXZdrjyFUX6ngiuXBevcXNEKCahwKgdSDRFzF7Cg=","RequestId":"ca90d525-d098-4a41-a7f2-d7606efd7a2f"}}

func encrypt_data() {
	credential := common.NewCredential(
		"IKIDqeKlxR8brszxFNSM9l9G7h15xc1C9uRk",
		"82jKisoY3RfoKp9JGxd5HExRf2Pckrre",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "kms.tencentcloudapi.com"
	client, _ := kms.NewClient(credential, "ap-hongkong", cpf)

	request := kms.NewEncryptRequest()

	request.KeyId = common.StringPtr("1fd7f96c-ea37-11ed-b68d-32c48ca45083") //your kms key_id,replace it
	request.Plaintext = common.StringPtr("MTIzNDU2Nzc4OQo=")                 //base64 of plaintext,replace it

	response, err := client.Encrypt(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("encrypt_data:%s\n", response.ToJsonString())
}

// {"Response":{"CiphertextBlob":"BYCUP6xePY1vS9u2xF2LoRGV4P7r7G5BwQ3rfaPvZaTZ7yU7IsgU8152hH8iPj8mn57MtKQrcX6fTNN9B11qPg==-k-fKVP3WIlGpg8m9LMW4jEkQ==-k-XihyJVRuO21xU9lJIBPxYQpX8XYeBQhXe2Dj2ewotFPcr+E1","KeyId":"1fd7f96c-ea37-11ed-b68d-32c48ca45083","RequestId":"b6be5337-408e-46bc-8705-dc4443a4225f"}}

func dencrypt_data() {
	credential := common.NewCredential(
		"IKIDqeKlxR8brszxFNSM9l9G7h15xc1C9uRk",
		"82jKisoY3RfoKp9JGxd5HExRf2Pckrre",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "kms.tencentcloudapi.com"
	client, _ := kms.NewClient(credential, "ap-hongkong", cpf)

	request := kms.NewDecryptRequest()

	//ciphertext,replace it
	request.CiphertextBlob = common.StringPtr("BYCUP6xePY1vS9u2xF2LoRGV4P7r7G5BwQ3rfaPvZaTZ7yU7IsgU8152hH8iPj8mn57MtKQrcX6fTNN9B11qPg==-k-fKVP3WIlGpg8m9LMW4jEkQ==-k-GkK58mChGcT4Re4N6NVDnh+VSlX7/9K396K27ZgaoUhF7VbO25xhA0h6VaoGSqseJeSzBX/yTapgpx40B9yzLG0vADlm+PCUnNbV6737NvmB16OefVxs9p/RB74F1NUL+60UPA==")

	response, err := client.Decrypt(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("dencrypt_data:%s\n", response.ToJsonString())
}

// {"Response":{"KeyId":"1fd7f96c-ea37-11ed-b68d-32c48ca45083","Plaintext":"MTIzNDU2Nzc4OQo=","RequestId":"5f1858b0-6fb0-49fa-8aac-a6368ed9a22c"}}
