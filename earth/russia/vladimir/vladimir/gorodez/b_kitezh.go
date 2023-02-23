package gorodez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基捷日KitezhBarony struct {
	feud.BaseBarony
}

var BKitezh基捷日 feud.Barony = &基捷日KitezhBarony{}

func init() {
	f := BKitezh基捷日.(*基捷日KitezhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kitezh",
		TitleName: "基捷日",
		TitleCode: "b_kitezh",
	}
}
