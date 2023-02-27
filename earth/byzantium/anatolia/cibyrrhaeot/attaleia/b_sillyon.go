package attaleia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希梁SillyonBarony struct {
	feud.BaseBarony
}

var BSillyon希梁 feud.Barony = &希梁SillyonBarony{}

func init() {
    f := BSillyon希梁.(*希梁SillyonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sillyon",
		TitleName: "希梁",
		TitleCode: "b_sillyon",
	}
}
