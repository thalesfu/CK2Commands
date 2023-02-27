package asturias_de_santillana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌迪亚莱斯堡CastrourdialesBarony struct {
	feud.BaseBarony
}

var BCastrourdiales乌迪亚莱斯堡 feud.Barony = &乌迪亚莱斯堡CastrourdialesBarony{}

func init() {
    f := BCastrourdiales乌迪亚莱斯堡.(*乌迪亚莱斯堡CastrourdialesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castrourdiales",
		TitleName: "乌迪亚莱斯堡",
		TitleCode: "b_castrourdiales",
	}
}
