package tarsos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朋多斯PendosisBarony struct {
	feud.BaseBarony
}

var BPendosis朋多斯 feud.Barony = &朋多斯PendosisBarony{}

func init() {
    f := BPendosis朋多斯.(*朋多斯PendosisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pendosis",
		TitleName: "朋多斯",
		TitleCode: "b_pendosis",
	}
}
