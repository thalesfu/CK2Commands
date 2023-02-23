package benevento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗里真托FrigentoBarony struct {
	feud.BaseBarony
}

var BFrigento弗里真托 feud.Barony = &弗里真托FrigentoBarony{}

func init() {
	f := BFrigento弗里真托.(*弗里真托FrigentoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "frigento",
		TitleName: "弗里真托",
		TitleCode: "b_frigento",
	}
}
