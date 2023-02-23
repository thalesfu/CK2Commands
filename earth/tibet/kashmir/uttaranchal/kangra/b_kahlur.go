package kangra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾罗尔KahlurBarony struct {
	feud.BaseBarony
}

var BKahlur贾罗尔 feud.Barony = &贾罗尔KahlurBarony{}

func init() {
	f := BKahlur贾罗尔.(*贾罗尔KahlurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kahlur",
		TitleName: "贾罗尔",
		TitleCode: "b_kahlur",
	}
}
