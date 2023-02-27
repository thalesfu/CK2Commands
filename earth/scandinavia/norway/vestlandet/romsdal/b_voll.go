package romsdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔VollBarony struct {
	feud.BaseBarony
}

var BVoll沃尔 feud.Barony = &沃尔VollBarony{}

func init() {
    f := BVoll沃尔.(*沃尔VollBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "voll",
		TitleName: "沃尔",
		TitleCode: "b_voll",
	}
}
