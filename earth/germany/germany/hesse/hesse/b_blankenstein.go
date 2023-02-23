package hesse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布兰肯施泰因BlankensteinBarony struct {
	feud.BaseBarony
}

var BBlankenstein布兰肯施泰因 feud.Barony = &布兰肯施泰因BlankensteinBarony{}

func init() {
	f := BBlankenstein布兰肯施泰因.(*布兰肯施泰因BlankensteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "blankenstein",
		TitleName: "布兰肯施泰因",
		TitleCode: "b_blankenstein",
	}
}
