package radha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗摩补罗菩阿梨阿Rampur_boaliaBarony struct {
	feud.BaseBarony
}

var BRampur_boalia罗摩补罗菩阿梨阿 feud.Barony = &罗摩补罗菩阿梨阿Rampur_boaliaBarony{}

func init() {
    f := BRampur_boalia罗摩补罗菩阿梨阿.(*罗摩补罗菩阿梨阿Rampur_boaliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rampur_boalia",
		TitleName: "罗摩补罗菩阿梨阿",
		TitleCode: "b_rampur_boalia",
	}
}
