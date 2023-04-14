package pb

import (
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

func TestPb() {
	bs := make([]byte, 0)
	var test TestData
	err := proto.Unmarshal(bs, &test)
	if err != nil {
		log.Info("fail")
	} else {
		v, err := json.Marshal(&test)
		if err != nil {
			fmt.Println("fail2")
		}
		fmt.Println(string(v))
	}
}
