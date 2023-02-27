package dobrzynska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基库乌KikotBarony struct {
	feud.BaseBarony
}

var BKikot基库乌 feud.Barony = &基库乌KikotBarony{}

func init() {
    f := BKikot基库乌.(*基库乌KikotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kikot",
		TitleName: "基库乌",
		TitleCode: "b_kikot",
	}
}
