package hedmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃兹库格EidskogBarony struct {
	feud.BaseBarony
}

var BEidskog埃兹库格 feud.Barony = &埃兹库格EidskogBarony{}

func init() {
	f := BEidskog埃兹库格.(*埃兹库格EidskogBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eidskog",
		TitleName: "埃兹库格",
		TitleCode: "b_eidskog",
	}
}
