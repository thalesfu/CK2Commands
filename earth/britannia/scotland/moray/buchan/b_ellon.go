package buchan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃伦EllonBarony struct {
	feud.BaseBarony
}

var BEllon埃伦 feud.Barony = &埃伦EllonBarony{}

func init() {
    f := BEllon埃伦.(*埃伦EllonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ellon",
		TitleName: "埃伦",
		TitleCode: "b_ellon",
	}
}
