package lincoln

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 林肯LincolnBarony struct {
	feud.BaseBarony
}

var BLincoln林肯 feud.Barony = &林肯LincolnBarony{}

func init() {
	f := BLincoln林肯.(*林肯LincolnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lincoln",
		TitleName: "林肯",
		TitleCode: "b_lincoln",
	}
}
