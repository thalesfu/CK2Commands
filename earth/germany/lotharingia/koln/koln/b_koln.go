package koln

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科隆KolnBarony struct {
	feud.BaseBarony
}

var BKoln科隆 feud.Barony = &科隆KolnBarony{}

func init() {
    f := BKoln科隆.(*科隆KolnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koln",
		TitleName: "科隆",
		TitleCode: "b_koln",
	}
}
