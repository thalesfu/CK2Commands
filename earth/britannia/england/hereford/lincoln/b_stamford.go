package lincoln

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯坦福StamfordBarony struct {
	feud.BaseBarony
}

var BStamford斯坦福 feud.Barony = &斯坦福StamfordBarony{}

func init() {
	f := BStamford斯坦福.(*斯坦福StamfordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stamford",
		TitleName: "斯坦福",
		TitleCode: "b_stamford",
	}
}
