package qazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尔马亚科KrmayakBarony struct {
	feud.BaseBarony
}

var BKrmayak克尔马亚科 feud.Barony = &克尔马亚科KrmayakBarony{}

func init() {
    f := BKrmayak克尔马亚科.(*克尔马亚科KrmayakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krmayak",
		TitleName: "克尔马亚科",
		TitleCode: "b_krmayak",
	}
}
