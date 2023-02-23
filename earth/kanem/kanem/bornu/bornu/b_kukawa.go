package bornu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库卡瓦KukawaBarony struct {
	feud.BaseBarony
}

var BKukawa库卡瓦 feud.Barony = &库卡瓦KukawaBarony{}

func init() {
	f := BKukawa库卡瓦.(*库卡瓦KukawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kukawa",
		TitleName: "库卡瓦",
		TitleCode: "b_kukawa",
	}
}
