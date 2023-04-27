package subject_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bingxueshuang/gnap/models/subject"
)

type SubInfo struct {
	SubIDs    []subject.ID `json:"sub_ids"`
	UpdatedAt time.Time    `json:"updated_at,omitempty"`
}

func Example() {
	userinfo, _ := subject.NewIDOpaque("J2G8G8O4AZ")
	lastupdated, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	user := SubInfo{
		[]subject.ID{userinfo},
		lastupdated,
	}
	data, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(data))
	// Output:
	// {
	//   "sub_ids": [
	//     {
	//       "format": "opaque",
	//       "id": "J2G8G8O4AZ"
	//     }
	//   ],
	//   "updated_at": "2006-01-02T15:04:05Z"
	// }
}
