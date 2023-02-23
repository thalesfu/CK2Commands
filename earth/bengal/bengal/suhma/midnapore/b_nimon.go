package midnapore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼牟NimonBarony struct {
	feud.BaseBarony
}

var BNimon尼牟 feud.Barony = &尼牟NimonBarony{}

func init() {
	f := BNimon尼牟.(*尼牟NimonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nimon",
		TitleName: "尼牟",
		TitleCode: "b_nimon",
	}
}
