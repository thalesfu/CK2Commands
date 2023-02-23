package orleans

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗勒里FleuryBarony struct {
	feud.BaseBarony
}

var BFleury弗勒里 feud.Barony = &弗勒里FleuryBarony{}

func init() {
	f := BFleury弗勒里.(*弗勒里FleuryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fleury",
		TitleName: "弗勒里",
		TitleCode: "b_fleury",
	}
}
