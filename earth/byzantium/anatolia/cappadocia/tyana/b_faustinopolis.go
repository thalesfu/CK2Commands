package tyana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福斯蒂诺波利斯FaustinopolisBarony struct {
	feud.BaseBarony
}

var BFaustinopolis福斯蒂诺波利斯 feud.Barony = &福斯蒂诺波利斯FaustinopolisBarony{}

func init() {
    f := BFaustinopolis福斯蒂诺波利斯.(*福斯蒂诺波利斯FaustinopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "faustinopolis",
		TitleName: "福斯蒂诺波利斯",
		TitleCode: "b_faustinopolis",
	}
}
