package scalovia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂尔吉特TilgitBarony struct {
	feud.BaseBarony
}

var BTilgit蒂尔吉特 feud.Barony = &蒂尔吉特TilgitBarony{}

func init() {
    f := BTilgit蒂尔吉特.(*蒂尔吉特TilgitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tilgit",
		TitleName: "蒂尔吉特",
		TitleCode: "b_tilgit",
	}
}
