package tell_bashir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布扎阿BuzaaBarony struct {
	feud.BaseBarony
}

var BBuzaa布扎阿 feud.Barony = &布扎阿BuzaaBarony{}

func init() {
    f := BBuzaa布扎阿.(*布扎阿BuzaaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buzaa",
		TitleName: "布扎阿",
		TitleCode: "b_buzaa",
	}
}
