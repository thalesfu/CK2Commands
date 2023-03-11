package cherven

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃伦VolynBarony struct {
	feud.BaseBarony
}

var BVolyn沃伦 feud.Barony = &沃伦VolynBarony{}

func init() {
    f := BVolyn沃伦.(*沃伦VolynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "volyn",
		TitleName: "沃伦",
		TitleCode: "b_volyn",
	}
}
