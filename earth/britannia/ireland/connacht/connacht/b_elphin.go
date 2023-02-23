package connacht

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔芬ElphinBarony struct {
	feud.BaseBarony
}

var BElphin埃尔芬 feud.Barony = &埃尔芬ElphinBarony{}

func init() {
	f := BElphin埃尔芬.(*埃尔芬ElphinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elphin",
		TitleName: "埃尔芬",
		TitleCode: "b_elphin",
	}
}
