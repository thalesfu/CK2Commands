package plzen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉托维KlatovyBarony struct {
	feud.BaseBarony
}

var BKlatovy克拉托维 feud.Barony = &克拉托维KlatovyBarony{}

func init() {
    f := BKlatovy克拉托维.(*克拉托维KlatovyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "klatovy",
		TitleName: "克拉托维",
		TitleCode: "b_klatovy",
	}
}
