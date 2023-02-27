package gurma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈洛戈HalogoBarony struct {
	feud.BaseBarony
}

var BHalogo哈洛戈 feud.Barony = &哈洛戈HalogoBarony{}

func init() {
    f := BHalogo哈洛戈.(*哈洛戈HalogoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "halogo",
		TitleName: "哈洛戈",
		TitleCode: "b_halogo",
	}
}
