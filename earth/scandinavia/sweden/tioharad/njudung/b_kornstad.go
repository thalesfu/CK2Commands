package njudung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科恩斯塔德KornstadBarony struct {
	feud.BaseBarony
}

var BKornstad科恩斯塔德 feud.Barony = &科恩斯塔德KornstadBarony{}

func init() {
    f := BKornstad科恩斯塔德.(*科恩斯塔德KornstadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kornstad",
		TitleName: "科恩斯塔德",
		TitleCode: "b_kornstad",
	}
}
