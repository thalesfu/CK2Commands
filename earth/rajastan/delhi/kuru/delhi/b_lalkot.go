package delhi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 剌罗拘吒LalkotBarony struct {
	feud.BaseBarony
}

var BLalkot剌罗拘吒 feud.Barony = &剌罗拘吒LalkotBarony{}

func init() {
	f := BLalkot剌罗拘吒.(*剌罗拘吒LalkotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lalkot",
		TitleName: "剌罗拘吒",
		TitleCode: "b_lalkot",
	}
}
