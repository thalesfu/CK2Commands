package trapani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯泰尔韦特拉诺CastelvertranoBarony struct {
	feud.BaseBarony
}

var BCastelvertrano卡斯泰尔韦特拉诺 feud.Barony = &卡斯泰尔韦特拉诺CastelvertranoBarony{}

func init() {
    f := BCastelvertrano卡斯泰尔韦特拉诺.(*卡斯泰尔韦特拉诺CastelvertranoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castelvertrano",
		TitleName: "卡斯泰尔韦特拉诺",
		TitleCode: "b_castelvertrano",
	}
}
