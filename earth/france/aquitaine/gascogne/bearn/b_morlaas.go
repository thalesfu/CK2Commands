package bearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫尔拉斯MorlaasBarony struct {
	feud.BaseBarony
}

var BMorlaas莫尔拉斯 feud.Barony = &莫尔拉斯MorlaasBarony{}

func init() {
	f := BMorlaas莫尔拉斯.(*莫尔拉斯MorlaasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "morlaas",
		TitleName: "莫尔拉斯",
		TitleCode: "b_morlaas",
	}
}
