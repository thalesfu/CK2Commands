package bolzano

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多拉迪兹DolladitzBarony struct {
	feud.BaseBarony
}

var BDolladitz多拉迪兹 feud.Barony = &多拉迪兹DolladitzBarony{}

func init() {
	f := BDolladitz多拉迪兹.(*多拉迪兹DolladitzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dolladitz",
		TitleName: "多拉迪兹",
		TitleCode: "b_dolladitz",
	}
}
