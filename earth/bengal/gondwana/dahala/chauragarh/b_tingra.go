package chauragarh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鼎揭罗TingraBarony struct {
	feud.BaseBarony
}

var BTingra鼎揭罗 feud.Barony = &鼎揭罗TingraBarony{}

func init() {
    f := BTingra鼎揭罗.(*鼎揭罗TingraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tingra",
		TitleName: "鼎揭罗",
		TitleCode: "b_tingra",
	}
}
