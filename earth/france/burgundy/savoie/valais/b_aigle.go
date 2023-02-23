package valais

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾格勒AigleBarony struct {
	feud.BaseBarony
}

var BAigle艾格勒 feud.Barony = &艾格勒AigleBarony{}

func init() {
	f := BAigle艾格勒.(*艾格勒AigleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aigle",
		TitleName: "艾格勒",
		TitleCode: "b_aigle",
	}
}
