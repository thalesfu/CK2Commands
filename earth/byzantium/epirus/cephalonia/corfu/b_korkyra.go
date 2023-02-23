package corfu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尔基拉KorkyraBarony struct {
	feud.BaseBarony
}

var BKorkyra克尔基拉 feud.Barony = &克尔基拉KorkyraBarony{}

func init() {
	f := BKorkyra克尔基拉.(*克尔基拉KorkyraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korkyra",
		TitleName: "克尔基拉",
		TitleCode: "b_korkyra",
	}
}
