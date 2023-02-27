package fitri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库卡KukaBarony struct {
	feud.BaseBarony
}

var BKuka库卡 feud.Barony = &库卡KukaBarony{}

func init() {
    f := BKuka库卡.(*库卡KukaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuka",
		TitleName: "库卡",
		TitleCode: "b_kuka",
	}
}
