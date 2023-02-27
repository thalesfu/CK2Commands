package tmutarakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特穆塔拉坎TmutarakanBarony struct {
	feud.BaseBarony
}

var BTmutarakan特穆塔拉坎 feud.Barony = &特穆塔拉坎TmutarakanBarony{}

func init() {
    f := BTmutarakan特穆塔拉坎.(*特穆塔拉坎TmutarakanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tmutarakan",
		TitleName: "特穆塔拉坎",
		TitleCode: "b_tmutarakan",
	}
}
