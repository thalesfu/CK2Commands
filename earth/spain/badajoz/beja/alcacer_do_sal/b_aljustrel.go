package alcacer_do_sal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔茹什特雷尔AljustrelBarony struct {
	feud.BaseBarony
}

var BAljustrel阿尔茹什特雷尔 feud.Barony = &阿尔茹什特雷尔AljustrelBarony{}

func init() {
    f := BAljustrel阿尔茹什特雷尔.(*阿尔茹什特雷尔AljustrelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aljustrel",
		TitleName: "阿尔茹什特雷尔",
		TitleCode: "b_aljustrel",
	}
}
