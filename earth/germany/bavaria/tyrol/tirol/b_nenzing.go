package tirol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 嫩青NenzingBarony struct {
	feud.BaseBarony
}

var BNenzing嫩青 feud.Barony = &嫩青NenzingBarony{}

func init() {
	f := BNenzing嫩青.(*嫩青NenzingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nenzing",
		TitleName: "嫩青",
		TitleCode: "b_nenzing",
	}
}
