package busaso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃里加博ErigavoBarony struct {
	feud.BaseBarony
}

var BErigavo埃里加博 feud.Barony = &埃里加博ErigavoBarony{}

func init() {
    f := BErigavo埃里加博.(*埃里加博ErigavoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erigavo",
		TitleName: "埃里加博",
		TitleCode: "b_erigavo",
	}
}
