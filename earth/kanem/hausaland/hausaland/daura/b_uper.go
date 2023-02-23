package daura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌珀UperBarony struct {
	feud.BaseBarony
}

var BUper乌珀 feud.Barony = &乌珀UperBarony{}

func init() {
	f := BUper乌珀.(*乌珀UperBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uper",
		TitleName: "乌珀",
		TitleCode: "b_uper",
	}
}
