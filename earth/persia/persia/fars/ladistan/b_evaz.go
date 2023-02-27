package ladistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃瓦兹EvazBarony struct {
	feud.BaseBarony
}

var BEvaz埃瓦兹 feud.Barony = &埃瓦兹EvazBarony{}

func init() {
    f := BEvaz埃瓦兹.(*埃瓦兹EvazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "evaz",
		TitleName: "埃瓦兹",
		TitleCode: "b_evaz",
	}
}
