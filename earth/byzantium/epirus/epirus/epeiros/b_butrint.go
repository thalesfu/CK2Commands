package epeiros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布特林特ButrintBarony struct {
	feud.BaseBarony
}

var BButrint布特林特 feud.Barony = &布特林特ButrintBarony{}

func init() {
	f := BButrint布特林特.(*布特林特ButrintBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "butrint",
		TitleName: "布特林特",
		TitleCode: "b_butrint",
	}
}
