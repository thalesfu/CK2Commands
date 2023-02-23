package marmaros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费尔舍维绍FelsovisoBarony struct {
	feud.BaseBarony
}

var BFelsoviso费尔舍维绍 feud.Barony = &费尔舍维绍FelsovisoBarony{}

func init() {
	f := BFelsoviso费尔舍维绍.(*费尔舍维绍FelsovisoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "felsoviso",
		TitleName: "费尔舍维绍",
		TitleCode: "b_felsoviso",
	}
}
