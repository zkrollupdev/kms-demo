// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

// [START kms_encrypt_symmetric]
import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"os"

	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// encryptSymmetric encrypts the input plaintext with the specified symmetric
// Cloud KMS key.
func encryptSymmetric(w io.Writer, name string, message string) []byte {
	// name := "projects/my-project/locations/us-east1/keyRings/my-key-ring/cryptoKeys/my-key"
	// message := "Sample message"

	// Create the client.
	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		println("a")
		return []byte{}

	}
	defer client.Close()

	// Convert the message into bytes. Cryptographic plaintexts and
	// ciphertexts are always byte arrays.
	plaintext := []byte(message)

	// Optional but recommended: Compute plaintext's CRC32C.
	crc32c := func(data []byte) uint32 {
		t := crc32.MakeTable(crc32.Castagnoli)
		return crc32.Checksum(data, t)
	}
	plaintextCRC32C := crc32c(plaintext)

	// Build the request.
	req := &kmspb.EncryptRequest{
		Name:            name,
		Plaintext:       plaintext,
		PlaintextCrc32C: wrapperspb.Int64(int64(plaintextCRC32C)),
	}

	// Call the API.
	result, err := client.Encrypt(ctx, req)
	if err != nil {
		println("b")
		fmt.Println(err)
		return []byte{}
	}

	// Optional, but recommended: perform integrity verification on result.
	// For more details on ensuring E2E in-transit integrity to and from Cloud KMS visit:
	// https://cloud.google.com/kms/docs/data-integrity-guidelines
	if result.VerifiedPlaintextCrc32C == false {
		println("c")
		return []byte{}
	}
	if int64(crc32c(result.Ciphertext)) != result.CiphertextCrc32C.Value {
		println("d")

		return []byte{}
	}
	// fmt.Printf(result.Ciphertext
	return result.Ciphertext
}

// decryptSymmetric will decrypt the input ciphertext bytes using the specified symmetric key.
func decryptSymmetric(w io.Writer, name string, ciphertext []byte) string {
	// name := "projects/my-project/locations/us-east1/keyRings/my-key-ring/cryptoKeys/my-key"
	// ciphertext := []byte("...")  // result of a symmetric encryption call

	// Create the client.
	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		return ""
	}
	defer client.Close()

	// Optional, but recommended: Compute ciphertext's CRC32C.
	crc32c := func(data []byte) uint32 {
		t := crc32.MakeTable(crc32.Castagnoli)
		return crc32.Checksum(data, t)
	}
	ciphertextCRC32C := crc32c(ciphertext)

	// Build the request.
	req := &kmspb.DecryptRequest{
		Name:             name,
		Ciphertext:       ciphertext,
		CiphertextCrc32C: wrapperspb.Int64(int64(ciphertextCRC32C)),
	}

	// Call the API.
	result, err := client.Decrypt(ctx, req)
	if err != nil {
		return ""
	}

	// Optional, but recommended: perform integrity verification on result.
	// For more details on ensuring E2E in-transit integrity to and from Cloud KMS visit:
	// https://cloud.google.com/kms/docs/data-integrity-guidelines
	if int64(crc32c(result.Plaintext)) != result.PlaintextCrc32C.Value {
		return ""
	}
	fmt.Fprintf(w, "Decrypted plaintext: %s", result.Plaintext)
	return string(result.Plaintext)
}

// [END kms_encrypt_symmetric]
func main() {
	var b bytes.Buffer
	name := "projects/secret-test-380802/locations/global/keyRings/jeason0405/cryptoKeys/quickstart"
	mingwen := "ea6c44ac03bff858b476bba40716402b03e41b8e97e276d1baec7c37d42484a0"

	miwen := encryptSymmetric(&b, name, mingwen)
	fmt.Println(miwen)

	fmt.Println(hex.EncodeToString(miwen))

	miwen2 := "0a240002caecda1f81732f723cba51ba65711bdae42143ba529abe33666602ca4bd34747710d126900870eaa7f8a5c2e9799c3ab6aabe4afa1df2e128b349b89ce2f3ad196b51c8431b11f445e1b382ba31294b94d09a1f9c73278402b316eef5bcd7a33decb04a0322e1a25ff84e41d92a8755d280efc8e95499c6c232f7d7e2c1ec4b47c7ef63deb904c36ac0e06483d"
	test123, _ := hex.DecodeString(miwen2)
	plain_text := decryptSymmetric(&b, name, test123)
	fmt.Println(plain_text)

	content, _ := os.ReadFile("./file")
	fmt.Println(hex.EncodeToString(content))

	plain_text_2 := decryptSymmetric(&b, name, content)
	fmt.Println(plain_text_2)

}
