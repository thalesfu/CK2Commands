package sticht

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾瑟尔斯泰因IjsselsteinBarony struct {
	feud.BaseBarony
}

var BIjsselstein艾瑟尔斯泰因 feud.Barony = &艾瑟尔斯泰因IjsselsteinBarony{}

func init() {
    f := BIjsselstein艾瑟尔斯泰因.(*艾瑟尔斯泰因IjsselsteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ijsselstein",
		TitleName: "艾瑟尔斯泰因",
		TitleCode: "b_ijsselstein",
	}
}
