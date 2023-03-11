package terebovl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 捷列博夫尔TerebovlBarony struct {
	feud.BaseBarony
}

var BTerebovl捷列博夫尔 feud.Barony = &捷列博夫尔TerebovlBarony{}

func init() {
    f := BTerebovl捷列博夫尔.(*捷列博夫尔TerebovlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "terebovl",
		TitleName: "捷列博夫尔",
		TitleCode: "b_terebovl",
	}
}
