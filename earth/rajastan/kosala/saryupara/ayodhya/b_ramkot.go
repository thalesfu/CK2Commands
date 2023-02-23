package ayodhya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗摩拘吒RamkotBarony struct {
	feud.BaseBarony
}

var BRamkot罗摩拘吒 feud.Barony = &罗摩拘吒RamkotBarony{}

func init() {
	f := BRamkot罗摩拘吒.(*罗摩拘吒RamkotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramkot",
		TitleName: "罗摩拘吒",
		TitleCode: "b_ramkot",
	}
}
