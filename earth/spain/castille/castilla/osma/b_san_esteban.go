package osma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣埃斯特万San_estebanBarony struct {
	feud.BaseBarony
}

var BSan_esteban圣埃斯特万 feud.Barony = &圣埃斯特万San_estebanBarony{}

func init() {
    f := BSan_esteban圣埃斯特万.(*圣埃斯特万San_estebanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "san_esteban",
		TitleName: "圣埃斯特万",
		TitleCode: "b_san_esteban",
	}
}
