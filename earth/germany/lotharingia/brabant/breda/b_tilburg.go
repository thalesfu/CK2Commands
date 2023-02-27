package breda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂尔堡TilburgBarony struct {
	feud.BaseBarony
}

var BTilburg蒂尔堡 feud.Barony = &蒂尔堡TilburgBarony{}

func init() {
    f := BTilburg蒂尔堡.(*蒂尔堡TilburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tilburg",
		TitleName: "蒂尔堡",
		TitleCode: "b_tilburg",
	}
}
