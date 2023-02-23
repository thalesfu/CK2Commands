package tummana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿摩罗建吒迦AmarkantakBarony struct {
	feud.BaseBarony
}

var BAmarkantak阿摩罗建吒迦 feud.Barony = &阿摩罗建吒迦AmarkantakBarony{}

func init() {
	f := BAmarkantak阿摩罗建吒迦.(*阿摩罗建吒迦AmarkantakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amarkantak",
		TitleName: "阿摩罗建吒迦",
		TitleCode: "b_amarkantak",
	}
}
