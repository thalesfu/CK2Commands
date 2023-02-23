package niebla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿罗谢ArocheBarony struct {
	feud.BaseBarony
}

var BAroche阿罗谢 feud.Barony = &阿罗谢ArocheBarony{}

func init() {
	f := BAroche阿罗谢.(*阿罗谢ArocheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aroche",
		TitleName: "阿罗谢",
		TitleCode: "b_aroche",
	}
}
