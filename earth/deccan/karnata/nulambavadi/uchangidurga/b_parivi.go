package uchangidurga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕里维PariviBarony struct {
	feud.BaseBarony
}

var BParivi帕里维 feud.Barony = &帕里维PariviBarony{}

func init() {
	f := BParivi帕里维.(*帕里维PariviBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "parivi",
		TitleName: "帕里维",
		TitleCode: "b_parivi",
	}
}
