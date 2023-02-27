package sarasvata_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 悉陀补罗SiddhapuraBarony struct {
	feud.BaseBarony
}

var BSiddhapura悉陀补罗 feud.Barony = &悉陀补罗SiddhapuraBarony{}

func init() {
    f := BSiddhapura悉陀补罗.(*悉陀补罗SiddhapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siddhapura",
		TitleName: "悉陀补罗",
		TitleCode: "b_siddhapura",
	}
}
