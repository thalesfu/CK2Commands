package marienburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马林堡MarienburgBarony struct {
	feud.BaseBarony
}

var BMarienburg马林堡 feud.Barony = &马林堡MarienburgBarony{}

func init() {
    f := BMarienburg马林堡.(*马林堡MarienburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marienburg",
		TitleName: "马林堡",
		TitleCode: "b_marienburg",
	}
}
