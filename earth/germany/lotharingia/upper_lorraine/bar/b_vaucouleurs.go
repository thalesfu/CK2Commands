package bar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃库勒尔VaucouleursBarony struct {
	feud.BaseBarony
}

var BVaucouleurs沃库勒尔 feud.Barony = &沃库勒尔VaucouleursBarony{}

func init() {
    f := BVaucouleurs沃库勒尔.(*沃库勒尔VaucouleursBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vaucouleurs",
		TitleName: "沃库勒尔",
		TitleCode: "b_vaucouleurs",
	}
}
