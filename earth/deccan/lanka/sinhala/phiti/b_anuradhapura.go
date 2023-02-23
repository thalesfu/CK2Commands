package phiti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿㝹罗陀补罗AnuradhapuraBarony struct {
	feud.BaseBarony
}

var BAnuradhapura阿㝹罗陀补罗 feud.Barony = &阿㝹罗陀补罗AnuradhapuraBarony{}

func init() {
	f := BAnuradhapura阿㝹罗陀补罗.(*阿㝹罗陀补罗AnuradhapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anuradhapura",
		TitleName: "阿㝹罗陀补罗",
		TitleCode: "b_anuradhapura",
	}
}
