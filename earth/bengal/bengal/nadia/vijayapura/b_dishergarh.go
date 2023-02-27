package vijayapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提室罗姞利呬DishergarhBarony struct {
	feud.BaseBarony
}

var BDishergarh提室罗姞利呬 feud.Barony = &提室罗姞利呬DishergarhBarony{}

func init() {
    f := BDishergarh提室罗姞利呬.(*提室罗姞利呬DishergarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dishergarh",
		TitleName: "提室罗姞利呬",
		TitleCode: "b_dishergarh",
	}
}
