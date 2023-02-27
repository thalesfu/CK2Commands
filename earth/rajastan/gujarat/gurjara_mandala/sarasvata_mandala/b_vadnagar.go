package sarasvata_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 跋那揭罗VadnagarBarony struct {
	feud.BaseBarony
}

var BVadnagar跋那揭罗 feud.Barony = &跋那揭罗VadnagarBarony{}

func init() {
    f := BVadnagar跋那揭罗.(*跋那揭罗VadnagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vadnagar",
		TitleName: "跋那揭罗",
		TitleCode: "b_vadnagar",
	}
}
