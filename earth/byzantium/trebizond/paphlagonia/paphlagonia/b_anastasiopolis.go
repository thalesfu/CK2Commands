package paphlagonia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿纳塔修波利斯AnastasiopolisBarony struct {
	feud.BaseBarony
}

var BAnastasiopolis阿纳塔修波利斯 feud.Barony = &阿纳塔修波利斯AnastasiopolisBarony{}

func init() {
    f := BAnastasiopolis阿纳塔修波利斯.(*阿纳塔修波利斯AnastasiopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anastasiopolis",
		TitleName: "阿纳塔修波利斯",
		TitleCode: "b_anastasiopolis",
	}
}
