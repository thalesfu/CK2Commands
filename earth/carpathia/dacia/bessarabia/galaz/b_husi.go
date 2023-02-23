package galaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡希HusiBarony struct {
	feud.BaseBarony
}

var BHusi胡希 feud.Barony = &胡希HusiBarony{}

func init() {
	f := BHusi胡希.(*胡希HusiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "husi",
		TitleName: "胡希",
		TitleCode: "b_husi",
	}
}
