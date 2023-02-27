package kolzum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科勒祖姆KolzumBarony struct {
	feud.BaseBarony
}

var BKolzum科勒祖姆 feud.Barony = &科勒祖姆KolzumBarony{}

func init() {
    f := BKolzum科勒祖姆.(*科勒祖姆KolzumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolzum",
		TitleName: "科勒祖姆",
		TitleCode: "b_kolzum",
	}
}
