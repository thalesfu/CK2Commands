package cremona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯泰尔戈弗雷多CastelgoffredoBarony struct {
	feud.BaseBarony
}

var BCastelgoffredo卡斯泰尔戈弗雷多 feud.Barony = &卡斯泰尔戈弗雷多CastelgoffredoBarony{}

func init() {
    f := BCastelgoffredo卡斯泰尔戈弗雷多.(*卡斯泰尔戈弗雷多CastelgoffredoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castelgoffredo",
		TitleName: "卡斯泰尔戈弗雷多",
		TitleCode: "b_castelgoffredo",
	}
}
