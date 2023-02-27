package khantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汉特ChantyBarony struct {
	feud.BaseBarony
}

var BChanty汉特 feud.Barony = &汉特ChantyBarony{}

func init() {
    f := BChanty汉特.(*汉特ChantyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chanty",
		TitleName: "汉特",
		TitleCode: "b_chanty",
	}
}
