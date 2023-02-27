package lorraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙尔佩涅CharpeigneBarony struct {
	feud.BaseBarony
}

var BCharpeigne沙尔佩涅 feud.Barony = &沙尔佩涅CharpeigneBarony{}

func init() {
    f := BCharpeigne沙尔佩涅.(*沙尔佩涅CharpeigneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charpeigne",
		TitleName: "沙尔佩涅",
		TitleCode: "b_charpeigne",
	}
}
