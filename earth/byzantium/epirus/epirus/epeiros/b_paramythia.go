package epeiros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕拉米西亚ParamythiaBarony struct {
	feud.BaseBarony
}

var BParamythia帕拉米西亚 feud.Barony = &帕拉米西亚ParamythiaBarony{}

func init() {
	f := BParamythia帕拉米西亚.(*帕拉米西亚ParamythiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paramythia",
		TitleName: "帕拉米西亚",
		TitleCode: "b_paramythia",
	}
}
