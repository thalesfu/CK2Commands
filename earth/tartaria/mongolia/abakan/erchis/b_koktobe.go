package erchis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科克托别KoktobeBarony struct {
	feud.BaseBarony
}

var BKoktobe科克托别 feud.Barony = &科克托别KoktobeBarony{}

func init() {
    f := BKoktobe科克托别.(*科克托别KoktobeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koktobe",
		TitleName: "科克托别",
		TitleCode: "b_koktobe",
	}
}
