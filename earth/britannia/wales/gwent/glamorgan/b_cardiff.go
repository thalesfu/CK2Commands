package glamorgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加的夫CardiffBarony struct {
	feud.BaseBarony
}

var BCardiff加的夫 feud.Barony = &加的夫CardiffBarony{}

func init() {
    f := BCardiff加的夫.(*加的夫CardiffBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cardiff",
		TitleName: "加的夫",
		TitleCode: "b_cardiff",
	}
}
