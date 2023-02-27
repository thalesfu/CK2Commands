package al_jawf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古里牙QurayyatBarony struct {
	feud.BaseBarony
}

var BQurayyat古里牙 feud.Barony = &古里牙QurayyatBarony{}

func init() {
    f := BQurayyat古里牙.(*古里牙QurayyatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qurayyat",
		TitleName: "古里牙",
		TitleCode: "b_qurayyat",
	}
}
