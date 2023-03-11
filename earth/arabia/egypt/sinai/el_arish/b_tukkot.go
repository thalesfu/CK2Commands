package el_arish

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图科特TukkotBarony struct {
	feud.BaseBarony
}

var BTukkot图科特 feud.Barony = &图科特TukkotBarony{}

func init() {
    f := BTukkot图科特.(*图科特TukkotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tukkot",
		TitleName: "图科特",
		TitleCode: "b_tukkot",
	}
}
