package arjin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 木孜塔格MuztaghBarony struct {
	feud.BaseBarony
}

var BMuztagh木孜塔格 feud.Barony = &木孜塔格MuztaghBarony{}

func init() {
	f := BMuztagh木孜塔格.(*木孜塔格MuztaghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muztagh",
		TitleName: "木孜塔格",
		TitleCode: "b_muztagh",
	}
}
