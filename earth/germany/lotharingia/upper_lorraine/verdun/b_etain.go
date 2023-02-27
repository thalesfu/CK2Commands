package verdun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃坦EtainBarony struct {
	feud.BaseBarony
}

var BEtain埃坦 feud.Barony = &埃坦EtainBarony{}

func init() {
    f := BEtain埃坦.(*埃坦EtainBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "etain",
		TitleName: "埃坦",
		TitleCode: "b_etain",
	}
}
