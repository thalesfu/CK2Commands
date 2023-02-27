package kubera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托克苏ToqsuBarony struct {
	feud.BaseBarony
}

var BToqsu托克苏 feud.Barony = &托克苏ToqsuBarony{}

func init() {
    f := BToqsu托克苏.(*托克苏ToqsuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "toqsu",
		TitleName: "托克苏",
		TitleCode: "b_toqsu",
	}
}
