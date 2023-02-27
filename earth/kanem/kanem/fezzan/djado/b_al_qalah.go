package djado

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖拉Al_qalahBarony struct {
	feud.BaseBarony
}

var BAl_qalah盖拉 feud.Barony = &盖拉Al_qalahBarony{}

func init() {
    f := BAl_qalah盖拉.(*盖拉Al_qalahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_qalah",
		TitleName: "盖拉",
		TitleCode: "b_al_qalah",
	}
}
