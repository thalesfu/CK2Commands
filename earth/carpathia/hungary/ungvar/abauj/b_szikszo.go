package abauj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡克索SzikszoBarony struct {
	feud.BaseBarony
}

var BSzikszo锡克索 feud.Barony = &锡克索SzikszoBarony{}

func init() {
	f := BSzikszo锡克索.(*锡克索SzikszoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szikszo",
		TitleName: "锡克索",
		TitleCode: "b_szikszo",
	}
}
