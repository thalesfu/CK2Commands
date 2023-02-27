package tmutarakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索乌贾克SujukqaleBarony struct {
	feud.BaseBarony
}

var BSujukqale索乌贾克 feud.Barony = &索乌贾克SujukqaleBarony{}

func init() {
    f := BSujukqale索乌贾克.(*索乌贾克SujukqaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sujukqale",
		TitleName: "索乌贾克",
		TitleCode: "b_sujukqale",
	}
}
