package angermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩舍尔兹维克OrnskoldsvikBarony struct {
	feud.BaseBarony
}

var BOrnskoldsvik恩舍尔兹维克 feud.Barony = &恩舍尔兹维克OrnskoldsvikBarony{}

func init() {
    f := BOrnskoldsvik恩舍尔兹维克.(*恩舍尔兹维克OrnskoldsvikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ornskoldsvik",
		TitleName: "恩舍尔兹维克",
		TitleCode: "b_ornskoldsvik",
	}
}
