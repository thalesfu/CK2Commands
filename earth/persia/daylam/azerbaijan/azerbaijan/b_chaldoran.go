package azerbaijan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 查尔多兰ChaldoranBarony struct {
	feud.BaseBarony
}

var BChaldoran查尔多兰 feud.Barony = &查尔多兰ChaldoranBarony{}

func init() {
	f := BChaldoran查尔多兰.(*查尔多兰ChaldoranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chaldoran",
		TitleName: "查尔多兰",
		TitleCode: "b_chaldoran",
	}
}
