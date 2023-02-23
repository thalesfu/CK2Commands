package wadai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔库莱KoulkouleBarony struct {
	feud.BaseBarony
}

var BKoulkoule库尔库莱 feud.Barony = &库尔库莱KoulkouleBarony{}

func init() {
	f := BKoulkoule库尔库莱.(*库尔库莱KoulkouleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koulkoule",
		TitleName: "库尔库莱",
		TitleCode: "b_koulkoule",
	}
}
