package marrakech

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代姆纳特DemnatBarony struct {
	feud.BaseBarony
}

var BDemnat代姆纳特 feud.Barony = &代姆纳特DemnatBarony{}

func init() {
	f := BDemnat代姆纳特.(*代姆纳特DemnatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "demnat",
		TitleName: "代姆纳特",
		TitleCode: "b_demnat",
	}
}
