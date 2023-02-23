package murcia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶克拉YeclaBarony struct {
	feud.BaseBarony
}

var BYecla耶克拉 feud.Barony = &耶克拉YeclaBarony{}

func init() {
	f := BYecla耶克拉.(*耶克拉YeclaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yecla",
		TitleName: "耶克拉",
		TitleCode: "b_yecla",
	}
}
