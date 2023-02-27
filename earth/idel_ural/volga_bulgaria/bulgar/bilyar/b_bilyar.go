package bilyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比利亚尔BilyarBarony struct {
	feud.BaseBarony
}

var BBilyar比利亚尔 feud.Barony = &比利亚尔BilyarBarony{}

func init() {
    f := BBilyar比利亚尔.(*比利亚尔BilyarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bilyar",
		TitleName: "比利亚尔",
		TitleCode: "b_bilyar",
	}
}
