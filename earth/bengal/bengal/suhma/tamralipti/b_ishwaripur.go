package tamralipti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊湿伐梨城IshwaripurBarony struct {
	feud.BaseBarony
}

var BIshwaripur伊湿伐梨城 feud.Barony = &伊湿伐梨城IshwaripurBarony{}

func init() {
	f := BIshwaripur伊湿伐梨城.(*伊湿伐梨城IshwaripurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ishwaripur",
		TitleName: "伊湿伐梨城",
		TitleCode: "b_ishwaripur",
	}
}
