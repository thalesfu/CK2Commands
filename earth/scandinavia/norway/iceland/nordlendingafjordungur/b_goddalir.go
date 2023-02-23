package nordlendingafjordungur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格兹达利尔GoddalirBarony struct {
	feud.BaseBarony
}

var BGoddalir格兹达利尔 feud.Barony = &格兹达利尔GoddalirBarony{}

func init() {
	f := BGoddalir格兹达利尔.(*格兹达利尔GoddalirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "goddalir",
		TitleName: "格兹达利尔",
		TitleCode: "b_goddalir",
	}
}
