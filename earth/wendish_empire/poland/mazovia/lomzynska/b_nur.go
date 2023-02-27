package lomzynska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努尔NurBarony struct {
	feud.BaseBarony
}

var BNur努尔 feud.Barony = &努尔NurBarony{}

func init() {
    f := BNur努尔.(*努尔NurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nur",
		TitleName: "努尔",
		TitleCode: "b_nur",
	}
}
