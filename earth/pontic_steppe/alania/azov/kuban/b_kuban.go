package kuban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库班KubanBarony struct {
	feud.BaseBarony
}

var BKuban库班 feud.Barony = &库班KubanBarony{}

func init() {
    f := BKuban库班.(*库班KubanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuban",
		TitleName: "库班",
		TitleCode: "b_kuban",
	}
}
