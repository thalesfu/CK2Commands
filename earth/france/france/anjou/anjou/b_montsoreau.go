package anjou

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙索罗MontsoreauBarony struct {
	feud.BaseBarony
}

var BMontsoreau蒙索罗 feud.Barony = &蒙索罗MontsoreauBarony{}

func init() {
	f := BMontsoreau蒙索罗.(*蒙索罗MontsoreauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montsoreau",
		TitleName: "蒙索罗",
		TitleCode: "b_montsoreau",
	}
}
