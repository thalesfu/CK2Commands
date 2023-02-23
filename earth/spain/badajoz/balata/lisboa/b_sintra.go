package lisboa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辛特拉SintraBarony struct {
	feud.BaseBarony
}

var BSintra辛特拉 feud.Barony = &辛特拉SintraBarony{}

func init() {
	f := BSintra辛特拉.(*辛特拉SintraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sintra",
		TitleName: "辛特拉",
		TitleCode: "b_sintra",
	}
}
