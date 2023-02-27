package sirjan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡尔詹SirjanBarony struct {
	feud.BaseBarony
}

var BSirjan锡尔詹 feud.Barony = &锡尔詹SirjanBarony{}

func init() {
    f := BSirjan锡尔詹.(*锡尔詹SirjanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sirjan",
		TitleName: "锡尔詹",
		TitleCode: "b_sirjan",
	}
}
