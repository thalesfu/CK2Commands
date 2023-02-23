package perche

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟农什SenonchesBarony struct {
	feud.BaseBarony
}

var BSenonches瑟农什 feud.Barony = &瑟农什SenonchesBarony{}

func init() {
	f := BSenonches瑟农什.(*瑟农什SenonchesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "senonches",
		TitleName: "瑟农什",
		TitleCode: "b_senonches",
	}
}
