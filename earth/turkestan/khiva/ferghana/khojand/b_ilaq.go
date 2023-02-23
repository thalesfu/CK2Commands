package khojand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 义剌克IlaqBarony struct {
	feud.BaseBarony
}

var BIlaq义剌克 feud.Barony = &义剌克IlaqBarony{}

func init() {
	f := BIlaq义剌克.(*义剌克IlaqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ilaq",
		TitleName: "义剌克",
		TitleCode: "b_ilaq",
	}
}
