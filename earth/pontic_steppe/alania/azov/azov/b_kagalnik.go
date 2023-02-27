package azov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡加利尼克KagalnikBarony struct {
	feud.BaseBarony
}

var BKagalnik卡加利尼克 feud.Barony = &卡加利尼克KagalnikBarony{}

func init() {
    f := BKagalnik卡加利尼克.(*卡加利尼克KagalnikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kagalnik",
		TitleName: "卡加利尼克",
		TitleCode: "b_kagalnik",
	}
}
