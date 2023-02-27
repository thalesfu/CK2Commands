package samoyeds

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔拉索瓦TarasovaBarony struct {
	feud.BaseBarony
}

var BTarasova塔拉索瓦 feud.Barony = &塔拉索瓦TarasovaBarony{}

func init() {
    f := BTarasova塔拉索瓦.(*塔拉索瓦TarasovaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarasova",
		TitleName: "塔拉索瓦",
		TitleCode: "b_tarasova",
	}
}
