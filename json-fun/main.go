package jsonfun

import (
	"encoding/json"
	"fmt"
)

func jsonEncode(data any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func main() {

}
