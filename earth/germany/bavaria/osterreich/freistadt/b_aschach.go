package freistadt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿沙赫AschachBarony struct {
	feud.BaseBarony
}

var BAschach阿沙赫 feud.Barony = &阿沙赫AschachBarony{}

func init() {
	f := BAschach阿沙赫.(*阿沙赫AschachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aschach",
		TitleName: "阿沙赫",
		TitleCode: "b_aschach",
	}
}
