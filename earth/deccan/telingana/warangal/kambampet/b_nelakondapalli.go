package kambampet

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泥罗军荼波梨NelakondapalliBarony struct {
	feud.BaseBarony
}

var BNelakondapalli泥罗军荼波梨 feud.Barony = &泥罗军荼波梨NelakondapalliBarony{}

func init() {
    f := BNelakondapalli泥罗军荼波梨.(*泥罗军荼波梨NelakondapalliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nelakondapalli",
		TitleName: "泥罗军荼波梨",
		TitleCode: "b_nelakondapalli",
	}
}
