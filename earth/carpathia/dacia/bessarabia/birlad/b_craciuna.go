package birlad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克勒丘内什蒂CraciunaBarony struct {
	feud.BaseBarony
}

var BCraciuna克勒丘内什蒂 feud.Barony = &克勒丘内什蒂CraciunaBarony{}

func init() {
    f := BCraciuna克勒丘内什蒂.(*克勒丘内什蒂CraciunaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "craciuna",
		TitleName: "克勒丘内什蒂",
		TitleCode: "b_craciuna",
	}
}
