package songhay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多莫塞拉DomosselaBarony struct {
	feud.BaseBarony
}

var BDomossela多莫塞拉 feud.Barony = &多莫塞拉DomosselaBarony{}

func init() {
	f := BDomossela多莫塞拉.(*多莫塞拉DomosselaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "domossela",
		TitleName: "多莫塞拉",
		TitleCode: "b_domossela",
	}
}
