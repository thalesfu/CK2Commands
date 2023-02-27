package teviotdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安嫩代尔PeeblesBarony struct {
	feud.BaseBarony
}

var BPeebles安嫩代尔 feud.Barony = &安嫩代尔PeeblesBarony{}

func init() {
    f := BPeebles安嫩代尔.(*安嫩代尔PeeblesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "peebles",
		TitleName: "安嫩代尔",
		TitleCode: "b_peebles",
	}
}
