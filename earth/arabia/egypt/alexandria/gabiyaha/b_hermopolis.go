package gabiyaha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫尔莫波利斯HermopolisBarony struct {
	feud.BaseBarony
}

var BHermopolis赫尔莫波利斯 feud.Barony = &赫尔莫波利斯HermopolisBarony{}

func init() {
    f := BHermopolis赫尔莫波利斯.(*赫尔莫波利斯HermopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hermopolis",
		TitleName: "赫尔莫波利斯",
		TitleCode: "b_hermopolis",
	}
}
