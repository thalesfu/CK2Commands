package kantor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索利科SolikoBarony struct {
	feud.BaseBarony
}

var BSoliko索利科 feud.Barony = &索利科SolikoBarony{}

func init() {
    f := BSoliko索利科.(*索利科SolikoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soliko",
		TitleName: "索利科",
		TitleCode: "b_soliko",
	}
}
