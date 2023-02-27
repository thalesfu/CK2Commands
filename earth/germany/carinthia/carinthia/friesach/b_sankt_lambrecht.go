package friesach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣兰布雷希特Sankt_lambrechtBarony struct {
	feud.BaseBarony
}

var BSankt_lambrecht圣兰布雷希特 feud.Barony = &圣兰布雷希特Sankt_lambrechtBarony{}

func init() {
    f := BSankt_lambrecht圣兰布雷希特.(*圣兰布雷希特Sankt_lambrechtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sankt_lambrecht",
		TitleName: "圣兰布雷希特",
		TitleCode: "b_sankt_lambrecht",
	}
}
