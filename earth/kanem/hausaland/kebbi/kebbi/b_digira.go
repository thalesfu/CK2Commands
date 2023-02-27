package kebbi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪吉拉DigiraBarony struct {
	feud.BaseBarony
}

var BDigira迪吉拉 feud.Barony = &迪吉拉DigiraBarony{}

func init() {
    f := BDigira迪吉拉.(*迪吉拉DigiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "digira",
		TitleName: "迪吉拉",
		TitleCode: "b_digira",
	}
}
