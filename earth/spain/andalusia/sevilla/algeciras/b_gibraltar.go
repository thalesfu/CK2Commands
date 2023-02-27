package algeciras

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 直布罗陀GibraltarBarony struct {
	feud.BaseBarony
}

var BGibraltar直布罗陀 feud.Barony = &直布罗陀GibraltarBarony{}

func init() {
    f := BGibraltar直布罗陀.(*直布罗陀GibraltarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gibraltar",
		TitleName: "直布罗陀",
		TitleCode: "b_gibraltar",
	}
}
