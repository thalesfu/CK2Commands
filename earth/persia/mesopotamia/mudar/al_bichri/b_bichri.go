package al_bichri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比基利BichriBarony struct {
	feud.BaseBarony
}

var BBichri比基利 feud.Barony = &比基利BichriBarony{}

func init() {
    f := BBichri比基利.(*比基利BichriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bichri",
		TitleName: "比基利",
		TitleCode: "b_bichri",
	}
}
