package models

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

var data = []byte("hello world")

func ExampleHashMethod() {
	var hash HashMethod
	_ = json.Unmarshal([]byte(`"sha-512"`), &hash)
	bytes := hash.Sum(data)
	fmt.Println(hex.EncodeToString(bytes))
	// Output:
	// 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
}
