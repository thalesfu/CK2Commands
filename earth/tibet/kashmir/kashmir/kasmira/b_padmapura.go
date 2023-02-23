package kasmira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 钵特摩补罗PadmapuraBarony struct {
	feud.BaseBarony
}

var BPadmapura钵特摩补罗 feud.Barony = &钵特摩补罗PadmapuraBarony{}

func init() {
	f := BPadmapura钵特摩补罗.(*钵特摩补罗PadmapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "padmapura",
		TitleName: "钵特摩补罗",
		TitleCode: "b_padmapura",
	}
}
