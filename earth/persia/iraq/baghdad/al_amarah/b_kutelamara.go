package al_amarah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库特阿玛拉KutelamaraBarony struct {
	feud.BaseBarony
}

var BKutelamara库特阿玛拉 feud.Barony = &库特阿玛拉KutelamaraBarony{}

func init() {
    f := BKutelamara库特阿玛拉.(*库特阿玛拉KutelamaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kutelamara",
		TitleName: "库特阿玛拉",
		TitleCode: "b_kutelamara",
	}
}
