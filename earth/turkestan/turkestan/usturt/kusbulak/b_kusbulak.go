package kusbulak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库斯布拉克KusbulakBarony struct {
	feud.BaseBarony
}

var BKusbulak库斯布拉克 feud.Barony = &库斯布拉克KusbulakBarony{}

func init() {
    f := BKusbulak库斯布拉克.(*库斯布拉克KusbulakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kusbulak",
		TitleName: "库斯布拉克",
		TitleCode: "b_kusbulak",
	}
}
