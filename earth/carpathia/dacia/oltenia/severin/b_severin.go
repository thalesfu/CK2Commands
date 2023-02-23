package severin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞韦林SeverinBarony struct {
	feud.BaseBarony
}

var BSeverin塞韦林 feud.Barony = &塞韦林SeverinBarony{}

func init() {
	f := BSeverin塞韦林.(*塞韦林SeverinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "severin",
		TitleName: "塞韦林",
		TitleCode: "b_severin",
	}
}
