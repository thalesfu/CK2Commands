package montbeliard

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙贝利亚尔MontbeliardBarony struct {
	feud.BaseBarony
}

var BMontbeliard蒙贝利亚尔 feud.Barony = &蒙贝利亚尔MontbeliardBarony{}

func init() {
    f := BMontbeliard蒙贝利亚尔.(*蒙贝利亚尔MontbeliardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montbeliard",
		TitleName: "蒙贝利亚尔",
		TitleCode: "b_montbeliard",
	}
}
