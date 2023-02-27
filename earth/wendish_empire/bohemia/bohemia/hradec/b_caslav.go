package hradec

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰斯拉夫CaslavBarony struct {
	feud.BaseBarony
}

var BCaslav恰斯拉夫 feud.Barony = &恰斯拉夫CaslavBarony{}

func init() {
    f := BCaslav恰斯拉夫.(*恰斯拉夫CaslavBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caslav",
		TitleName: "恰斯拉夫",
		TitleCode: "b_caslav",
	}
}
