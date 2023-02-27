package goa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦补罗姞利呬KapurgarhBarony struct {
	feud.BaseBarony
}

var BKapurgarh迦补罗姞利呬 feud.Barony = &迦补罗姞利呬KapurgarhBarony{}

func init() {
    f := BKapurgarh迦补罗姞利呬.(*迦补罗姞利呬KapurgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kapurgarh",
		TitleName: "迦补罗姞利呬",
		TitleCode: "b_kapurgarh",
	}
}
