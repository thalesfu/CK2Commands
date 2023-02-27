package vestfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿兰达拉ArendallBarony struct {
	feud.BaseBarony
}

var BArendall阿兰达拉 feud.Barony = &阿兰达拉ArendallBarony{}

func init() {
    f := BArendall阿兰达拉.(*阿兰达拉ArendallBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arendall",
		TitleName: "阿兰达拉",
		TitleCode: "b_arendall",
	}
}
