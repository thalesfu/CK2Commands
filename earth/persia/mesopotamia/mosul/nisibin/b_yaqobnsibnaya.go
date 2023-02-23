package nisibin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚喀尼斯纳亚YaqobnsibnayaBarony struct {
	feud.BaseBarony
}

var BYaqobnsibnaya亚喀尼斯纳亚 feud.Barony = &亚喀尼斯纳亚YaqobnsibnayaBarony{}

func init() {
	f := BYaqobnsibnaya亚喀尼斯纳亚.(*亚喀尼斯纳亚YaqobnsibnayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yaqobnsibnaya",
		TitleName: "亚喀尼斯纳亚",
		TitleCode: "b_yaqobnsibnaya",
	}
}
