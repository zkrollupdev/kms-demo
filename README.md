## try kms func demo of tencent cloud
    go run kms-tencent.go

## try kms func demo of google cloud
    go run kms-gcp.go


## code example for tencent kms env

### 1.generate CMK
```
    func main() {

        credential := common.NewCredential(
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
            fmt.Printf("An API error has returned: %s", err)
            return
        }
        if err != nil {
            panic(err)
        }
        fmt.Printf("%s", response.ToJsonString())

    }
```

### 2.generate data encrypt key 
```
func main() {
	credential := common.NewCredential(
		"IKIDqeKlxR8brszxFNSM9l9G7h15xc1C9uRk",
		"82jKisoY3RfoKp9JGxd5HExRf2Pckrre",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "kms.tencentcloudapi.com"
	client, _ := kms.NewClient(credential, "ap-hongkong", cpf)

	request := kms.NewGenerateDataKeyRequest()

	request.KeyId = common.StringPtr("1fd7f96c-ea37-11ed-b68d-32c48ca45083")
	request.KeySpec = common.StringPtr("AES_256")
	
	response, err := client.GenerateDataKey(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
```

### 3.encrypt data using the key above
```
func main() {
	credential := common.NewCredential(
		"IKIDqeKlxR8brszxFNSM9l9G7h15xc1C9uRk",
		"82jKisoY3RfoKp9JGxd5HExRf2Pckrre",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "kms.tencentcloudapi.com"
	client, _ := kms.NewClient(credential, "ap-hongkong", cpf)

	request := kms.NewEncryptRequest()

	request.KeyId = common.StringPtr("1fd7f96c-ea37-11ed-b68d-32c48ca45083")
	request.Plaintext = common.StringPtr("MTIzNDU2Nzc4OQo=")

	response, err := client.Encrypt(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
```
request.Plaintext:
    must be the base64 format,like:MTIzNDU2Nzc4OQo=,you can use the linux comnmand base64 to transfer some msg to it,for example:
    echo '1234567789' | base64


### 4.decrypt data using the key above
```
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
```
response:
    this function will give you the base64 format of plaintext,you can transfer it to plaintext msg using the command like this:
    echo 'MTIzNDU2Nzc4OQo' | base64 -d





