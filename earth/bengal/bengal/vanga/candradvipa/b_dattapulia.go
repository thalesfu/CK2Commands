package candradvipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀怛多补梨DattapuliaBarony struct {
	feud.BaseBarony
}

var BDattapulia陀怛多补梨 feud.Barony = &陀怛多补梨DattapuliaBarony{}

func init() {
	f := BDattapulia陀怛多补梨.(*陀怛多补梨DattapuliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dattapulia",
		TitleName: "陀怛多补梨",
		TitleCode: "b_dattapulia",
	}
}
