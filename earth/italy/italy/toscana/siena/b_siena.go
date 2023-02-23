package siena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡耶纳SienaBarony struct {
	feud.BaseBarony
}

var BSiena锡耶纳 feud.Barony = &锡耶纳SienaBarony{}

func init() {
	f := BSiena锡耶纳.(*锡耶纳SienaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siena",
		TitleName: "锡耶纳",
		TitleCode: "b_siena",
	}
}
