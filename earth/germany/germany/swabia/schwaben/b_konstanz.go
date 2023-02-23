package schwaben

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 康斯坦茨KonstanzBarony struct {
	feud.BaseBarony
}

var BKonstanz康斯坦茨 feud.Barony = &康斯坦茨KonstanzBarony{}

func init() {
	f := BKonstanz康斯坦茨.(*康斯坦茨KonstanzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "konstanz",
		TitleName: "康斯坦茨",
		TitleCode: "b_konstanz",
	}
}
