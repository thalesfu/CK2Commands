package kirkuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尔库克KirkukBarony struct {
	feud.BaseBarony
}

var BKirkuk基尔库克 feud.Barony = &基尔库克KirkukBarony{}

func init() {
    f := BKirkuk基尔库克.(*基尔库克KirkukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirkuk",
		TitleName: "基尔库克",
		TitleCode: "b_kirkuk",
	}
}
