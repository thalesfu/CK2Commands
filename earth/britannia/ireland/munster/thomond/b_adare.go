package thomond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿代尔AdareBarony struct {
	feud.BaseBarony
}

var BAdare阿代尔 feud.Barony = &阿代尔AdareBarony{}

func init() {
	f := BAdare阿代尔.(*阿代尔AdareBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adare",
		TitleName: "阿代尔",
		TitleCode: "b_adare",
	}
}
