package leicester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂克希尔TickhillBarony struct {
	feud.BaseBarony
}

var BTickhill蒂克希尔 feud.Barony = &蒂克希尔TickhillBarony{}

func init() {
	f := BTickhill蒂克希尔.(*蒂克希尔TickhillBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tickhill",
		TitleName: "蒂克希尔",
		TitleCode: "b_tickhill",
	}
}
