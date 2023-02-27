package ouled_nail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾格瓦特Al_aghwatBarony struct {
	feud.BaseBarony
}

var BAl_aghwat艾格瓦特 feud.Barony = &艾格瓦特Al_aghwatBarony{}

func init() {
    f := BAl_aghwat艾格瓦特.(*艾格瓦特Al_aghwatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_aghwat",
		TitleName: "艾格瓦特",
		TitleCode: "b_al_aghwat",
	}
}
