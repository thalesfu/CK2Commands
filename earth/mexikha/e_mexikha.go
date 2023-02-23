package mexikha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MexikhaEmpire interface {
	feud.Empire
}

type 阿兹特克帝国MexikhaEmpire struct {
	feud.BaseEmpire
}

var EMexikha阿兹特克帝国 MexikhaEmpire = &阿兹特克帝国MexikhaEmpire{}

func init() {
	f := EMexikha阿兹特克帝国.(*阿兹特克帝国MexikhaEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "mexikha",
		TitleName: "阿兹特克帝国",
		TitleCode: "e_mexikha",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
