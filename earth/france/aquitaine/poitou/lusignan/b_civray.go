package lusignan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡夫赖CivrayBarony struct {
	feud.BaseBarony
}

var BCivray锡夫赖 feud.Barony = &锡夫赖CivrayBarony{}

func init() {
	f := BCivray锡夫赖.(*锡夫赖CivrayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "civray",
		TitleName: "锡夫赖",
		TitleCode: "b_civray",
	}
}
