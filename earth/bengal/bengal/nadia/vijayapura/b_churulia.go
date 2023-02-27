package vijayapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 周楼梨耶ChuruliaBarony struct {
	feud.BaseBarony
}

var BChurulia周楼梨耶 feud.Barony = &周楼梨耶ChuruliaBarony{}

func init() {
    f := BChurulia周楼梨耶.(*周楼梨耶ChuruliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "churulia",
		TitleName: "周楼梨耶",
		TitleCode: "b_churulia",
	}
}
