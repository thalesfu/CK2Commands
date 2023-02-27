package semender

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库穆赫KumukhBarony struct {
	feud.BaseBarony
}

var BKumukh库穆赫 feud.Barony = &库穆赫KumukhBarony{}

func init() {
    f := BKumukh库穆赫.(*库穆赫KumukhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kumukh",
		TitleName: "库穆赫",
		TitleCode: "b_kumukh",
	}
}
