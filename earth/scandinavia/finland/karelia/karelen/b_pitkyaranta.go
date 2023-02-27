package karelen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮特基亚兰塔PitkyarantaBarony struct {
	feud.BaseBarony
}

var BPitkyaranta皮特基亚兰塔 feud.Barony = &皮特基亚兰塔PitkyarantaBarony{}

func init() {
    f := BPitkyaranta皮特基亚兰塔.(*皮特基亚兰塔PitkyarantaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pitkyaranta",
		TitleName: "皮特基亚兰塔",
		TitleCode: "b_pitkyaranta",
	}
}
