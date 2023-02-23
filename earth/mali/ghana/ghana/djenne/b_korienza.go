package djenne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科里安扎KorienzaBarony struct {
	feud.BaseBarony
}

var BKorienza科里安扎 feud.Barony = &科里安扎KorienzaBarony{}

func init() {
	f := BKorienza科里安扎.(*科里安扎KorienzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korienza",
		TitleName: "科里安扎",
		TitleCode: "b_korienza",
	}
}
