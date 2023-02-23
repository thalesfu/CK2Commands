package smolensk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅特里佐沃NetrizovoBarony struct {
	feud.BaseBarony
}

var BNetrizovo涅特里佐沃 feud.Barony = &涅特里佐沃NetrizovoBarony{}

func init() {
	f := BNetrizovo涅特里佐沃.(*涅特里佐沃NetrizovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "netrizovo",
		TitleName: "涅特里佐沃",
		TitleCode: "b_netrizovo",
	}
}
