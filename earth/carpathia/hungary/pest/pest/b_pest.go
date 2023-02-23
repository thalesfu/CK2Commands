package pest

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩斯PestBarony struct {
	feud.BaseBarony
}

var BPest佩斯 feud.Barony = &佩斯PestBarony{}

func init() {
	f := BPest佩斯.(*佩斯PestBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pest",
		TitleName: "佩斯",
		TitleCode: "b_pest",
	}
}
