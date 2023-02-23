package burhanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃伦多尔ErandolBarony struct {
	feud.BaseBarony
}

var BErandol埃伦多尔 feud.Barony = &埃伦多尔ErandolBarony{}

func init() {
	f := BErandol埃伦多尔.(*埃伦多尔ErandolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erandol",
		TitleName: "埃伦多尔",
		TitleCode: "b_erandol",
	}
}
