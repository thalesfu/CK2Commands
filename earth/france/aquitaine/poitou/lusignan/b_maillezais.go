package lusignan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马耶赛MaillezaisBarony struct {
	feud.BaseBarony
}

var BMaillezais马耶赛 feud.Barony = &马耶赛MaillezaisBarony{}

func init() {
    f := BMaillezais马耶赛.(*马耶赛MaillezaisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maillezais",
		TitleName: "马耶赛",
		TitleCode: "b_maillezais",
	}
}
