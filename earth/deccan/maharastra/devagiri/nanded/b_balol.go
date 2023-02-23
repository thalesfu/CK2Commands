package nanded

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆卢罗BalolBarony struct {
	feud.BaseBarony
}

var BBalol婆卢罗 feud.Barony = &婆卢罗BalolBarony{}

func init() {
	f := BBalol婆卢罗.(*婆卢罗BalolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balol",
		TitleName: "婆卢罗",
		TitleCode: "b_balol",
	}
}
