package byzantion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣索菲亚HagiasophiaBarony struct {
	feud.BaseBarony
}

var BHagiasophia圣索菲亚 feud.Barony = &圣索菲亚HagiasophiaBarony{}

func init() {
    f := BHagiasophia圣索菲亚.(*圣索菲亚HagiasophiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hagiasophia",
		TitleName: "圣索菲亚",
		TitleCode: "b_hagiasophia",
	}
}
