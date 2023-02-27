package arborea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣朱斯塔Santa_giustaBarony struct {
	feud.BaseBarony
}

var BSanta_giusta圣朱斯塔 feud.Barony = &圣朱斯塔Santa_giustaBarony{}

func init() {
    f := BSanta_giusta圣朱斯塔.(*圣朱斯塔Santa_giustaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "santa_giusta",
		TitleName: "圣朱斯塔",
		TitleCode: "b_santa_giusta",
	}
}
