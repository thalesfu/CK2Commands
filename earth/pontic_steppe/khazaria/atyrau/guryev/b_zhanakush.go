package guryev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎那库什ZhanakushBarony struct {
	feud.BaseBarony
}

var BZhanakush扎那库什 feud.Barony = &扎那库什ZhanakushBarony{}

func init() {
    f := BZhanakush扎那库什.(*扎那库什ZhanakushBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhanakush",
		TitleName: "扎那库什",
		TitleCode: "b_zhanakush",
	}
}
