package hayya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海亚HayyaBarony struct {
	feud.BaseBarony
}

var BHayya海亚 feud.Barony = &海亚HayyaBarony{}

func init() {
	f := BHayya海亚.(*海亚HayyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hayya",
		TitleName: "海亚",
		TitleCode: "b_hayya",
	}
}
