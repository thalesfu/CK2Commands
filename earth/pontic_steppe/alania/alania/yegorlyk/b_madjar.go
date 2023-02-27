package yegorlyk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马佳尔MadjarBarony struct {
	feud.BaseBarony
}

var BMadjar马佳尔 feud.Barony = &马佳尔MadjarBarony{}

func init() {
    f := BMadjar马佳尔.(*马佳尔MadjarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madjar",
		TitleName: "马佳尔",
		TitleCode: "b_madjar",
	}
}
