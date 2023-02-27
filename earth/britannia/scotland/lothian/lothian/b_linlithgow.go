package lothian

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 林利斯戈LinlithgowBarony struct {
	feud.BaseBarony
}

var BLinlithgow林利斯戈 feud.Barony = &林利斯戈LinlithgowBarony{}

func init() {
    f := BLinlithgow林利斯戈.(*林利斯戈LinlithgowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "linlithgow",
		TitleName: "林利斯戈",
		TitleCode: "b_linlithgow",
	}
}
