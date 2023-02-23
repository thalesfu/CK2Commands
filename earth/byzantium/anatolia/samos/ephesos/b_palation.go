package ephesos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕拉提翁PalationBarony struct {
	feud.BaseBarony
}

var BPalation帕拉提翁 feud.Barony = &帕拉提翁PalationBarony{}

func init() {
	f := BPalation帕拉提翁.(*帕拉提翁PalationBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palation",
		TitleName: "帕拉提翁",
		TitleCode: "b_palation",
	}
}
