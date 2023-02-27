package east_karelia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索罗卡SorokaBarony struct {
	feud.BaseBarony
}

var BSoroka索罗卡 feud.Barony = &索罗卡SorokaBarony{}

func init() {
    f := BSoroka索罗卡.(*索罗卡SorokaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soroka",
		TitleName: "索罗卡",
		TitleCode: "b_soroka",
	}
}
