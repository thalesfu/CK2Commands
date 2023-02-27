package khokh_serkh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多罗DuruuBarony struct {
	feud.BaseBarony
}

var BDuruu多罗 feud.Barony = &多罗DuruuBarony{}

func init() {
    f := BDuruu多罗.(*多罗DuruuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "duruu",
		TitleName: "多罗",
		TitleCode: "b_duruu",
	}
}
