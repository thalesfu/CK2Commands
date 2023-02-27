package domazlice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍亨富尔特HohenfurthBarony struct {
	feud.BaseBarony
}

var BHohenfurth霍亨富尔特 feud.Barony = &霍亨富尔特HohenfurthBarony{}

func init() {
    f := BHohenfurth霍亨富尔特.(*霍亨富尔特HohenfurthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hohenfurth",
		TitleName: "霍亨富尔特",
		TitleCode: "b_hohenfurth",
	}
}
