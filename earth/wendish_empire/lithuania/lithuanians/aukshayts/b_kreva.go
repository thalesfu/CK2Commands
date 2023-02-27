package aukshayts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克列沃KrevaBarony struct {
	feud.BaseBarony
}

var BKreva克列沃 feud.Barony = &克列沃KrevaBarony{}

func init() {
    f := BKreva克列沃.(*克列沃KrevaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kreva",
		TitleName: "克列沃",
		TitleCode: "b_kreva",
	}
}
