package al_hasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佛达FodaBarony struct {
	feud.BaseBarony
}

var BFoda佛达 feud.Barony = &佛达FodaBarony{}

func init() {
    f := BFoda佛达.(*佛达FodaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "foda",
		TitleName: "佛达",
		TitleCode: "b_foda",
	}
}
