package amaravati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马切尔拉MacherlaBarony struct {
	feud.BaseBarony
}

var BMacherla马切尔拉 feud.Barony = &马切尔拉MacherlaBarony{}

func init() {
	f := BMacherla马切尔拉.(*马切尔拉MacherlaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "macherla",
		TitleName: "马切尔拉",
		TitleCode: "b_macherla",
	}
}
