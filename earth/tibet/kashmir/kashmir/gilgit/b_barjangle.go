package gilgit

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔姜戈BarjangleBarony struct {
	feud.BaseBarony
}

var BBarjangle巴尔姜戈 feud.Barony = &巴尔姜戈BarjangleBarony{}

func init() {
    f := BBarjangle巴尔姜戈.(*巴尔姜戈BarjangleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barjangle",
		TitleName: "巴尔姜戈",
		TitleCode: "b_barjangle",
	}
}
