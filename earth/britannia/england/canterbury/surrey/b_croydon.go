package surrey

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克罗伊登CroydonBarony struct {
	feud.BaseBarony
}

var BCroydon克罗伊登 feud.Barony = &克罗伊登CroydonBarony{}

func init() {
	f := BCroydon克罗伊登.(*克罗伊登CroydonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "croydon",
		TitleName: "克罗伊登",
		TitleCode: "b_croydon",
	}
}
