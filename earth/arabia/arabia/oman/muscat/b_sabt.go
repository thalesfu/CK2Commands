package muscat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞卜特SabtBarony struct {
	feud.BaseBarony
}

var BSabt塞卜特 feud.Barony = &塞卜特SabtBarony{}

func init() {
	f := BSabt塞卜特.(*塞卜特SabtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sabt",
		TitleName: "塞卜特",
		TitleCode: "b_sabt",
	}
}
