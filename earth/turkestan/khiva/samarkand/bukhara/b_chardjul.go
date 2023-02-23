package bukhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 查尔朱ChardjulBarony struct {
	feud.BaseBarony
}

var BChardjul查尔朱 feud.Barony = &查尔朱ChardjulBarony{}

func init() {
	f := BChardjul查尔朱.(*查尔朱ChardjulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chardjul",
		TitleName: "查尔朱",
		TitleCode: "b_chardjul",
	}
}
