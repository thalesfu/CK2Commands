package balkh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达利韦尔津捷佩DalverzintepeBarony struct {
	feud.BaseBarony
}

var BDalverzintepe达利韦尔津捷佩 feud.Barony = &达利韦尔津捷佩DalverzintepeBarony{}

func init() {
    f := BDalverzintepe达利韦尔津捷佩.(*达利韦尔津捷佩DalverzintepeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dalverzintepe",
		TitleName: "达利韦尔津捷佩",
		TitleCode: "b_dalverzintepe",
	}
}
