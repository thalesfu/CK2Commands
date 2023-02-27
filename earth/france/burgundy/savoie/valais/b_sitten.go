package valais

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡滕SittenBarony struct {
	feud.BaseBarony
}

var BSitten锡滕 feud.Barony = &锡滕SittenBarony{}

func init() {
    f := BSitten锡滕.(*锡滕SittenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sitten",
		TitleName: "锡滕",
		TitleCode: "b_sitten",
	}
}
