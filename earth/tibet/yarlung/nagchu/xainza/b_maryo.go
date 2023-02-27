package xainza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马跃MaryoBarony struct {
	feud.BaseBarony
}

var BMaryo马跃 feud.Barony = &马跃MaryoBarony{}

func init() {
    f := BMaryo马跃.(*马跃MaryoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maryo",
		TitleName: "马跃",
		TitleCode: "b_maryo",
	}
}
