package nishapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古昌QuchanBarony struct {
	feud.BaseBarony
}

var BQuchan古昌 feud.Barony = &古昌QuchanBarony{}

func init() {
    f := BQuchan古昌.(*古昌QuchanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "quchan",
		TitleName: "古昌",
		TitleCode: "b_quchan",
	}
}
