package songhay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂拉贝里TillaberiBarony struct {
	feud.BaseBarony
}

var BTillaberi蒂拉贝里 feud.Barony = &蒂拉贝里TillaberiBarony{}

func init() {
    f := BTillaberi蒂拉贝里.(*蒂拉贝里TillaberiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tillaberi",
		TitleName: "蒂拉贝里",
		TitleCode: "b_tillaberi",
	}
}
