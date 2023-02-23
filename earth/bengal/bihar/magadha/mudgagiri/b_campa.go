package mudgagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞻波CampaBarony struct {
	feud.BaseBarony
}

var BCampa瞻波 feud.Barony = &瞻波CampaBarony{}

func init() {
	f := BCampa瞻波.(*瞻波CampaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "campa",
		TitleName: "瞻波",
		TitleCode: "b_campa",
	}
}
