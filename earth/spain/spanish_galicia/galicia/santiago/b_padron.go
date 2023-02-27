package santiago

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕德龙PadronBarony struct {
	feud.BaseBarony
}

var BPadron帕德龙 feud.Barony = &帕德龙PadronBarony{}

func init() {
    f := BPadron帕德龙.(*帕德龙PadronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "padron",
		TitleName: "帕德龙",
		TitleCode: "b_padron",
	}
}
