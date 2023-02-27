package siena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙特普尔恰诺MontepulcianoBarony struct {
	feud.BaseBarony
}

var BMontepulciano蒙特普尔恰诺 feud.Barony = &蒙特普尔恰诺MontepulcianoBarony{}

func init() {
    f := BMontepulciano蒙特普尔恰诺.(*蒙特普尔恰诺MontepulcianoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montepulciano",
		TitleName: "蒙特普尔恰诺",
		TitleCode: "b_montepulciano",
	}
}
