package auvergne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙特费朗MontferrandBarony struct {
	feud.BaseBarony
}

var BMontferrand蒙特费朗 feud.Barony = &蒙特费朗MontferrandBarony{}

func init() {
    f := BMontferrand蒙特费朗.(*蒙特费朗MontferrandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montferrand",
		TitleName: "蒙特费朗",
		TitleCode: "b_montferrand",
	}
}
