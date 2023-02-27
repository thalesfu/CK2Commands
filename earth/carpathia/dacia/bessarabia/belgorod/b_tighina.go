package belgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂吉纳TighinaBarony struct {
	feud.BaseBarony
}

var BTighina蒂吉纳 feud.Barony = &蒂吉纳TighinaBarony{}

func init() {
    f := BTighina蒂吉纳.(*蒂吉纳TighinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tighina",
		TitleName: "蒂吉纳",
		TitleCode: "b_tighina",
	}
}
