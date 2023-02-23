package jharkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆娑摩BhasmaBarony struct {
	feud.BaseBarony
}

var BBhasma婆娑摩 feud.Barony = &婆娑摩BhasmaBarony{}

func init() {
	f := BBhasma婆娑摩.(*婆娑摩BhasmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhasma",
		TitleName: "婆娑摩",
		TitleCode: "b_bhasma",
	}
}
