package baden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔夫CalwBarony struct {
	feud.BaseBarony
}

var BCalw卡尔夫 feud.Barony = &卡尔夫CalwBarony{}

func init() {
	f := BCalw卡尔夫.(*卡尔夫CalwBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calw",
		TitleName: "卡尔夫",
		TitleCode: "b_calw",
	}
}
