package hardsyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔堡UlburgBarony struct {
	feud.BaseBarony
}

var BUlburg乌尔堡 feud.Barony = &乌尔堡UlburgBarony{}

func init() {
	f := BUlburg乌尔堡.(*乌尔堡UlburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulburg",
		TitleName: "乌尔堡",
		TitleCode: "b_ulburg",
	}
}
