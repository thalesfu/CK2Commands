package viken

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厄克勒OckeroBarony struct {
	feud.BaseBarony
}

var BOckero厄克勒 feud.Barony = &厄克勒OckeroBarony{}

func init() {
    f := BOckero厄克勒.(*厄克勒OckeroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ockero",
		TitleName: "厄克勒",
		TitleCode: "b_ockero",
	}
}
