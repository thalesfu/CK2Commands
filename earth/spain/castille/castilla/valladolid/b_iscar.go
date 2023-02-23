package valladolid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯卡尔IscarBarony struct {
	feud.BaseBarony
}

var BIscar伊斯卡尔 feud.Barony = &伊斯卡尔IscarBarony{}

func init() {
	f := BIscar伊斯卡尔.(*伊斯卡尔IscarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iscar",
		TitleName: "伊斯卡尔",
		TitleCode: "b_iscar",
	}
}
