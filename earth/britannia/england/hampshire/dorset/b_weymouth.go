package dorset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦茅斯WeymouthBarony struct {
	feud.BaseBarony
}

var BWeymouth韦茅斯 feud.Barony = &韦茅斯WeymouthBarony{}

func init() {
    f := BWeymouth韦茅斯.(*韦茅斯WeymouthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "weymouth",
		TitleName: "韦茅斯",
		TitleCode: "b_weymouth",
	}
}
