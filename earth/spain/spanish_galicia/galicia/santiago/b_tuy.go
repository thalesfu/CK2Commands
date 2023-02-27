package santiago

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图伊TuyBarony struct {
	feud.BaseBarony
}

var BTuy图伊 feud.Barony = &图伊TuyBarony{}

func init() {
    f := BTuy图伊.(*图伊TuyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tuy",
		TitleName: "图伊",
		TitleCode: "b_tuy",
	}
}
