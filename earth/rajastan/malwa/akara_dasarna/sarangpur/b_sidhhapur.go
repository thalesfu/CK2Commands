package sarangpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 悉陀补罗SidhhapurBarony struct {
	feud.BaseBarony
}

var BSidhhapur悉陀补罗 feud.Barony = &悉陀补罗SidhhapurBarony{}

func init() {
    f := BSidhhapur悉陀补罗.(*悉陀补罗SidhhapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sidhhapur",
		TitleName: "悉陀补罗",
		TitleCode: "b_sidhhapur",
	}
}
