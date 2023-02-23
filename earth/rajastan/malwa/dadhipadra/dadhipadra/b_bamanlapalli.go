package dadhipadra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆曼罗波梨BamanlapalliBarony struct {
	feud.BaseBarony
}

var BBamanlapalli婆曼罗波梨 feud.Barony = &婆曼罗波梨BamanlapalliBarony{}

func init() {
	f := BBamanlapalli婆曼罗波梨.(*婆曼罗波梨BamanlapalliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bamanlapalli",
		TitleName: "婆曼罗波梨",
		TitleCode: "b_bamanlapalli",
	}
}
