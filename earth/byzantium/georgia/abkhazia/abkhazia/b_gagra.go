package abkhazia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加格拉GagraBarony struct {
	feud.BaseBarony
}

var BGagra加格拉 feud.Barony = &加格拉GagraBarony{}

func init() {
	f := BGagra加格拉.(*加格拉GagraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gagra",
		TitleName: "加格拉",
		TitleCode: "b_gagra",
	}
}
