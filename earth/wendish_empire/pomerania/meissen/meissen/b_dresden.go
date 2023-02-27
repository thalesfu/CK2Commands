package meissen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德累斯顿DresdenBarony struct {
	feud.BaseBarony
}

var BDresden德累斯顿 feud.Barony = &德累斯顿DresdenBarony{}

func init() {
    f := BDresden德累斯顿.(*德累斯顿DresdenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dresden",
		TitleName: "德累斯顿",
		TitleCode: "b_dresden",
	}
}
