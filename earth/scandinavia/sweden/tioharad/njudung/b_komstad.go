package njudung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科姆斯塔德KomstadBarony struct {
	feud.BaseBarony
}

var BKomstad科姆斯塔德 feud.Barony = &科姆斯塔德KomstadBarony{}

func init() {
    f := BKomstad科姆斯塔德.(*科姆斯塔德KomstadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "komstad",
		TitleName: "科姆斯塔德",
		TitleCode: "b_komstad",
	}
}
