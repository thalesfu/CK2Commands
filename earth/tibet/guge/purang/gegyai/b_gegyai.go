package gegyai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 革吉GegyaiBarony struct {
	feud.BaseBarony
}

var BGegyai革吉 feud.Barony = &革吉GegyaiBarony{}

func init() {
	f := BGegyai革吉.(*革吉GegyaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gegyai",
		TitleName: "革吉",
		TitleCode: "b_gegyai",
	}
}
