package corsica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔维Calvi_bBarony struct {
	feud.BaseBarony
}

var BCalvi_b卡尔维 feud.Barony = &卡尔维Calvi_bBarony{}

func init() {
    f := BCalvi_b卡尔维.(*卡尔维Calvi_bBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calvi_b",
		TitleName: "卡尔维",
		TitleCode: "b_calvi_b",
	}
}
