package joensuu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊洛曼齐IlomantsiBarony struct {
	feud.BaseBarony
}

var BIlomantsi伊洛曼齐 feud.Barony = &伊洛曼齐IlomantsiBarony{}

func init() {
    f := BIlomantsi伊洛曼齐.(*伊洛曼齐IlomantsiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ilomantsi",
		TitleName: "伊洛曼齐",
		TitleCode: "b_ilomantsi",
	}
}
