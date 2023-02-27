package santiago

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣地亚哥SantiagoBarony struct {
	feud.BaseBarony
}

var BSantiago圣地亚哥 feud.Barony = &圣地亚哥SantiagoBarony{}

func init() {
    f := BSantiago圣地亚哥.(*圣地亚哥SantiagoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "santiago",
		TitleName: "圣地亚哥",
		TitleCode: "b_santiago",
	}
}
