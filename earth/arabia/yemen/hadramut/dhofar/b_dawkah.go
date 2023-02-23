package dhofar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 道卡DawkahBarony struct {
	feud.BaseBarony
}

var BDawkah道卡 feud.Barony = &道卡DawkahBarony{}

func init() {
	f := BDawkah道卡.(*道卡DawkahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dawkah",
		TitleName: "道卡",
		TitleCode: "b_dawkah",
	}
}
