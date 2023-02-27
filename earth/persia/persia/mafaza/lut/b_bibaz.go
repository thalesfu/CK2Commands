package lut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比巴茨BibazBarony struct {
	feud.BaseBarony
}

var BBibaz比巴茨 feud.Barony = &比巴茨BibazBarony{}

func init() {
    f := BBibaz比巴茨.(*比巴茨BibazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bibaz",
		TitleName: "比巴茨",
		TitleCode: "b_bibaz",
	}
}
