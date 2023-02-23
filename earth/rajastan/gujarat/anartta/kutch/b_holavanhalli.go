package kutch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呼罗槃诃梨HolavanhalliBarony struct {
	feud.BaseBarony
}

var BHolavanhalli呼罗槃诃梨 feud.Barony = &呼罗槃诃梨HolavanhalliBarony{}

func init() {
	f := BHolavanhalli呼罗槃诃梨.(*呼罗槃诃梨HolavanhalliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "holavanhalli",
		TitleName: "呼罗槃诃梨",
		TitleCode: "b_holavanhalli",
	}
}
