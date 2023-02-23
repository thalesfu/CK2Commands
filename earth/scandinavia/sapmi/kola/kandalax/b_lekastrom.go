package kandalax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列卡斯特罗姆LekastromBarony struct {
	feud.BaseBarony
}

var BLekastrom列卡斯特罗姆 feud.Barony = &列卡斯特罗姆LekastromBarony{}

func init() {
	f := BLekastrom列卡斯特罗姆.(*列卡斯特罗姆LekastromBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lekastrom",
		TitleName: "列卡斯特罗姆",
		TitleCode: "b_lekastrom",
	}
}
