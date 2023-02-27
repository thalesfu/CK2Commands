package kudalasangama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鸠努尔KuknurBarony struct {
	feud.BaseBarony
}

var BKuknur鸠努尔 feud.Barony = &鸠努尔KuknurBarony{}

func init() {
    f := BKuknur鸠努尔.(*鸠努尔KuknurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuknur",
		TitleName: "鸠努尔",
		TitleCode: "b_kuknur",
	}
}
