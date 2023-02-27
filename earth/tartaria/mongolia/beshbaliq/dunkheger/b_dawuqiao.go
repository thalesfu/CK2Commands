package dunkheger

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大坞桥DawuqiaoBarony struct {
	feud.BaseBarony
}

var BDawuqiao大坞桥 feud.Barony = &大坞桥DawuqiaoBarony{}

func init() {
    f := BDawuqiao大坞桥.(*大坞桥DawuqiaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dawuqiao",
		TitleName: "大坞桥",
		TitleCode: "b_dawuqiao",
	}
}
