package hum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅敦MedunBarony struct {
	feud.BaseBarony
}

var BMedun梅敦 feud.Barony = &梅敦MedunBarony{}

func init() {
	f := BMedun梅敦.(*梅敦MedunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medun",
		TitleName: "梅敦",
		TitleCode: "b_medun",
	}
}
