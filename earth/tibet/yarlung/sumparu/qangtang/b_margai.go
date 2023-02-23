package qangtang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛尔盖MargaiBarony struct {
	feud.BaseBarony
}

var BMargai玛尔盖 feud.Barony = &玛尔盖MargaiBarony{}

func init() {
	f := BMargai玛尔盖.(*玛尔盖MargaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "margai",
		TitleName: "玛尔盖",
		TitleCode: "b_margai",
	}
}
