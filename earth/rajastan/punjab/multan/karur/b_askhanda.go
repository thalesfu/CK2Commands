package karur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯坎达AskhandaBarony struct {
	feud.BaseBarony
}

var BAskhanda阿斯坎达 feud.Barony = &阿斯坎达AskhandaBarony{}

func init() {
	f := BAskhanda阿斯坎达.(*阿斯坎达AskhandaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "askhanda",
		TitleName: "阿斯坎达",
		TitleCode: "b_askhanda",
	}
}
