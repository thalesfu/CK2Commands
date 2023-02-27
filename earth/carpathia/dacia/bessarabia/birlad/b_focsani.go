package birlad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福克沙尼FocsaniBarony struct {
	feud.BaseBarony
}

var BFocsani福克沙尼 feud.Barony = &福克沙尼FocsaniBarony{}

func init() {
    f := BFocsani福克沙尼.(*福克沙尼FocsaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "focsani",
		TitleName: "福克沙尼",
		TitleCode: "b_focsani",
	}
}
