package narim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯基兹AskizBarony struct {
	feud.BaseBarony
}

var BAskiz阿斯基兹 feud.Barony = &阿斯基兹AskizBarony{}

func init() {
    f := BAskiz阿斯基兹.(*阿斯基兹AskizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "askiz",
		TitleName: "阿斯基兹",
		TitleCode: "b_askiz",
	}
}
