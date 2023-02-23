package urgench

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎雷库利KanlykulBarony struct {
	feud.BaseBarony
}

var BKanlykul坎雷库利 feud.Barony = &坎雷库利KanlykulBarony{}

func init() {
	f := BKanlykul坎雷库利.(*坎雷库利KanlykulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kanlykul",
		TitleName: "坎雷库利",
		TitleCode: "b_kanlykul",
	}
}
