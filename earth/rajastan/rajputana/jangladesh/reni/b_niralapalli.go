package reni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼罗罗波梨NiralapalliBarony struct {
	feud.BaseBarony
}

var BNiralapalli尼罗罗波梨 feud.Barony = &尼罗罗波梨NiralapalliBarony{}

func init() {
	f := BNiralapalli尼罗罗波梨.(*尼罗罗波梨NiralapalliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "niralapalli",
		TitleName: "尼罗罗波梨",
		TitleCode: "b_niralapalli",
	}
}
