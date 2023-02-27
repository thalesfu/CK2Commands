package burtasy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅切特纳亚MechetnayaBarony struct {
	feud.BaseBarony
}

var BMechetnaya梅切特纳亚 feud.Barony = &梅切特纳亚MechetnayaBarony{}

func init() {
    f := BMechetnaya梅切特纳亚.(*梅切特纳亚MechetnayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mechetnaya",
		TitleName: "梅切特纳亚",
		TitleCode: "b_mechetnaya",
	}
}
