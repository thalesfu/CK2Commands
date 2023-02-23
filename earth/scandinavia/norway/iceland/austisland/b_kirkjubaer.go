package austisland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尔丘拜尔KirkjubaerBarony struct {
	feud.BaseBarony
}

var BKirkjubaer基尔丘拜尔 feud.Barony = &基尔丘拜尔KirkjubaerBarony{}

func init() {
	f := BKirkjubaer基尔丘拜尔.(*基尔丘拜尔KirkjubaerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirkjubaer",
		TitleName: "基尔丘拜尔",
		TitleCode: "b_kirkjubaer",
	}
}
