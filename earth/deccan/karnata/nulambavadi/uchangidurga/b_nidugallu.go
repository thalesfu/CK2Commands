package uchangidurga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼杜加卢NidugalluBarony struct {
	feud.BaseBarony
}

var BNidugallu尼杜加卢 feud.Barony = &尼杜加卢NidugalluBarony{}

func init() {
    f := BNidugallu尼杜加卢.(*尼杜加卢NidugalluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nidugallu",
		TitleName: "尼杜加卢",
		TitleCode: "b_nidugallu",
	}
}
