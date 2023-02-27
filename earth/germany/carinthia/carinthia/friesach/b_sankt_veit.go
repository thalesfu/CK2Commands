package friesach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣法伊特Sankt_veitBarony struct {
	feud.BaseBarony
}

var BSankt_veit圣法伊特 feud.Barony = &圣法伊特Sankt_veitBarony{}

func init() {
    f := BSankt_veit圣法伊特.(*圣法伊特Sankt_veitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sankt_veit",
		TitleName: "圣法伊特",
		TitleCode: "b_sankt_veit",
	}
}
