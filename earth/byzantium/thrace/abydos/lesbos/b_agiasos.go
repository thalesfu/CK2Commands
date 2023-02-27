package lesbos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿亚索斯AgiasosBarony struct {
	feud.BaseBarony
}

var BAgiasos阿亚索斯 feud.Barony = &阿亚索斯AgiasosBarony{}

func init() {
    f := BAgiasos阿亚索斯.(*阿亚索斯AgiasosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agiasos",
		TitleName: "阿亚索斯",
		TitleCode: "b_agiasos",
	}
}
