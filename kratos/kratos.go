package kratos

import (
	"encoding/json"
	"fmt"
	"github.com/artstylecode/artcoding-go/reflectutils"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

func FillParams(fillType interface{}, c *bm.Context) interface{} {

	testObj := make(map[string]interface{}, 0)
	testObj["name"] = "test1"
	testObj["id"] = 12
	testObj["age"] = 33
	paramStr, _ := json.Marshal(c.Params)
	fmt.Printf("debug val:%s\n", string(paramStr))
	targetInstance := reflectutils.NewInstance(fillType)
	json.Unmarshal(paramStr, targetInstance)
	fmt.Sprintf("value:%v", targetInstance)
	return targetInstance

}
