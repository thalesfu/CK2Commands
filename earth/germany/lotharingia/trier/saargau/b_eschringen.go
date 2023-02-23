package saargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃施林根EschringenBarony struct {
	feud.BaseBarony
}

var BEschringen埃施林根 feud.Barony = &埃施林根EschringenBarony{}

func init() {
	f := BEschringen埃施林根.(*埃施林根EschringenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eschringen",
		TitleName: "埃施林根",
		TitleCode: "b_eschringen",
	}
}
