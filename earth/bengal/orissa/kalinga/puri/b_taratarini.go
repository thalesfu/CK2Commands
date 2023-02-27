package puri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔拉塔里尼TaratariniBarony struct {
	feud.BaseBarony
}

var BTaratarini塔拉塔里尼 feud.Barony = &塔拉塔里尼TaratariniBarony{}

func init() {
    f := BTaratarini塔拉塔里尼.(*塔拉塔里尼TaratariniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taratarini",
		TitleName: "塔拉塔里尼",
		TitleCode: "b_taratarini",
	}
}
