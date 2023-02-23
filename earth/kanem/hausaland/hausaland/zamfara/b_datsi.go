package zamfara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达齐DatsiBarony struct {
	feud.BaseBarony
}

var BDatsi达齐 feud.Barony = &达齐DatsiBarony{}

func init() {
	f := BDatsi达齐.(*达齐DatsiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "datsi",
		TitleName: "达齐",
		TitleCode: "b_datsi",
	}
}
