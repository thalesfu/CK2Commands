package toulouse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米雷MuretBarony struct {
	feud.BaseBarony
}

var BMuret米雷 feud.Barony = &米雷MuretBarony{}

func init() {
    f := BMuret米雷.(*米雷MuretBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muret",
		TitleName: "米雷",
		TitleCode: "b_muret",
	}
}
