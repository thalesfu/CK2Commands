package korchev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博斯普鲁斯BosphorosBarony struct {
	feud.BaseBarony
}

var BBosphoros博斯普鲁斯 feud.Barony = &博斯普鲁斯BosphorosBarony{}

func init() {
    f := BBosphoros博斯普鲁斯.(*博斯普鲁斯BosphorosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bosphoros",
		TitleName: "博斯普鲁斯",
		TitleCode: "b_bosphoros",
	}
}
