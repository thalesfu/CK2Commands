package gemer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛雄茨LosoncBarony struct {
	feud.BaseBarony
}

var BLosonc洛雄茨 feud.Barony = &洛雄茨LosoncBarony{}

func init() {
    f := BLosonc洛雄茨.(*洛雄茨LosoncBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "losonc",
		TitleName: "洛雄茨",
		TitleCode: "b_losonc",
	}
}
