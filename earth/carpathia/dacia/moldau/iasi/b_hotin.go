package iasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍京HotinBarony struct {
	feud.BaseBarony
}

var BHotin霍京 feud.Barony = &霍京HotinBarony{}

func init() {
    f := BHotin霍京.(*霍京HotinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hotin",
		TitleName: "霍京",
		TitleCode: "b_hotin",
	}
}
