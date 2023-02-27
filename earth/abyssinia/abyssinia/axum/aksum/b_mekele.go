package aksum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 默克莱MekeleBarony struct {
	feud.BaseBarony
}

var BMekele默克莱 feud.Barony = &默克莱MekeleBarony{}

func init() {
    f := BMekele默克莱.(*默克莱MekeleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mekele",
		TitleName: "默克莱",
		TitleCode: "b_mekele",
	}
}
