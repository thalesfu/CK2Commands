package herakleia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波利斯PolisBarony struct {
	feud.BaseBarony
}

var BPolis波利斯 feud.Barony = &波利斯PolisBarony{}

func init() {
	f := BPolis波利斯.(*波利斯PolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "polis",
		TitleName: "波利斯",
		TitleCode: "b_polis",
	}
}
