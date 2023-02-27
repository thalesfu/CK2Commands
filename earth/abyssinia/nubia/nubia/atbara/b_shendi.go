package atbara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尚迪ShendiBarony struct {
	feud.BaseBarony
}

var BShendi尚迪 feud.Barony = &尚迪ShendiBarony{}

func init() {
    f := BShendi尚迪.(*尚迪ShendiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shendi",
		TitleName: "尚迪",
		TitleCode: "b_shendi",
	}
}
