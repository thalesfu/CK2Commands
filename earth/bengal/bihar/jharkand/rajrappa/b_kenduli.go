package rajrappa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艺那头梨KenduliBarony struct {
	feud.BaseBarony
}

var BKenduli艺那头梨 feud.Barony = &艺那头梨KenduliBarony{}

func init() {
    f := BKenduli艺那头梨.(*艺那头梨KenduliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kenduli",
		TitleName: "艺那头梨",
		TitleCode: "b_kenduli",
	}
}
