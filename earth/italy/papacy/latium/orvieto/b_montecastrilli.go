package orvieto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙特卡斯特里利MontecastrilliBarony struct {
	feud.BaseBarony
}

var BMontecastrilli蒙特卡斯特里利 feud.Barony = &蒙特卡斯特里利MontecastrilliBarony{}

func init() {
    f := BMontecastrilli蒙特卡斯特里利.(*蒙特卡斯特里利MontecastrilliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montecastrilli",
		TitleName: "蒙特卡斯特里利",
		TitleCode: "b_montecastrilli",
	}
}
