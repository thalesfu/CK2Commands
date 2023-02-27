package cairo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫利奥波利斯HeliopolisBarony struct {
	feud.BaseBarony
}

var BHeliopolis赫利奥波利斯 feud.Barony = &赫利奥波利斯HeliopolisBarony{}

func init() {
    f := BHeliopolis赫利奥波利斯.(*赫利奥波利斯HeliopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "heliopolis",
		TitleName: "赫利奥波利斯",
		TitleCode: "b_heliopolis",
	}
}
