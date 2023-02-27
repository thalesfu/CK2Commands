package sanaa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨那SanaaBarony struct {
	feud.BaseBarony
}

var BSanaa萨那 feud.Barony = &萨那SanaaBarony{}

func init() {
    f := BSanaa萨那.(*萨那SanaaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanaa",
		TitleName: "萨那",
		TitleCode: "b_sanaa",
	}
}
