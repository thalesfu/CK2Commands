package sikkim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加欣GyalshingBarony struct {
	feud.BaseBarony
}

var BGyalshing加欣 feud.Barony = &加欣GyalshingBarony{}

func init() {
	f := BGyalshing加欣.(*加欣GyalshingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gyalshing",
		TitleName: "加欣",
		TitleCode: "b_gyalshing",
	}
}
