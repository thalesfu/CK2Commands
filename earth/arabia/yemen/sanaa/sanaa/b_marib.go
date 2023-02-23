package sanaa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马里卜MaribBarony struct {
	feud.BaseBarony
}

var BMarib马里卜 feud.Barony = &马里卜MaribBarony{}

func init() {
	f := BMarib马里卜.(*马里卜MaribBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marib",
		TitleName: "马里卜",
		TitleCode: "b_marib",
	}
}
