package montpellier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾格莫尔特AiguesmortesBarony struct {
	feud.BaseBarony
}

var BAiguesmortes艾格莫尔特 feud.Barony = &艾格莫尔特AiguesmortesBarony{}

func init() {
    f := BAiguesmortes艾格莫尔特.(*艾格莫尔特AiguesmortesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aiguesmortes",
		TitleName: "艾格莫尔特",
		TitleCode: "b_aiguesmortes",
	}
}
