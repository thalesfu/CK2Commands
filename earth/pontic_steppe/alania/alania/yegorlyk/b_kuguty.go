package yegorlyk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库古特KugutyBarony struct {
	feud.BaseBarony
}

var BKuguty库古特 feud.Barony = &库古特KugutyBarony{}

func init() {
    f := BKuguty库古特.(*库古特KugutyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuguty",
		TitleName: "库古特",
		TitleCode: "b_kuguty",
	}
}
