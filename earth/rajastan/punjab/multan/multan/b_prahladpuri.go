package multan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 钵罗曷罗陀补利PrahladpuriBarony struct {
	feud.BaseBarony
}

var BPrahladpuri钵罗曷罗陀补利 feud.Barony = &钵罗曷罗陀补利PrahladpuriBarony{}

func init() {
    f := BPrahladpuri钵罗曷罗陀补利.(*钵罗曷罗陀补利PrahladpuriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prahladpuri",
		TitleName: "钵罗曷罗陀补利",
		TitleCode: "b_prahladpuri",
	}
}
