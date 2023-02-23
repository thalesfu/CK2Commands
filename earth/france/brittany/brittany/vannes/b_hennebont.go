package vannes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃讷邦HennebontBarony struct {
	feud.BaseBarony
}

var BHennebont埃讷邦 feud.Barony = &埃讷邦HennebontBarony{}

func init() {
	f := BHennebont埃讷邦.(*埃讷邦HennebontBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hennebont",
		TitleName: "埃讷邦",
		TitleCode: "b_hennebont",
	}
}
