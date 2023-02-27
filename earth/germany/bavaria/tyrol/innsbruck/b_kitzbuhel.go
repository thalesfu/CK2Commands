package innsbruck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基茨比厄尔KitzbuhelBarony struct {
	feud.BaseBarony
}

var BKitzbuhel基茨比厄尔 feud.Barony = &基茨比厄尔KitzbuhelBarony{}

func init() {
    f := BKitzbuhel基茨比厄尔.(*基茨比厄尔KitzbuhelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kitzbuhel",
		TitleName: "基茨比厄尔",
		TitleCode: "b_kitzbuhel",
	}
}
