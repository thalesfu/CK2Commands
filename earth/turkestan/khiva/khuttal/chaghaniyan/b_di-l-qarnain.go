package chaghaniyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 左勒盖尔奈英Di_l_qarnainBarony struct {
	feud.BaseBarony
}

var BDi_l_qarnain左勒盖尔奈英 feud.Barony = &左勒盖尔奈英Di_l_qarnainBarony{}

func init() {
    f := BDi_l_qarnain左勒盖尔奈英.(*左勒盖尔奈英Di_l_qarnainBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "di_l_qarnain",
		TitleName: "左勒盖尔奈英",
		TitleCode: "b_di-l-qarnain",
	}
}
