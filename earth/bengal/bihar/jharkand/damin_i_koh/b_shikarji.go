package damin_i_koh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尸佉罗时ShikarjiBarony struct {
	feud.BaseBarony
}

var BShikarji尸佉罗时 feud.Barony = &尸佉罗时ShikarjiBarony{}

func init() {
    f := BShikarji尸佉罗时.(*尸佉罗时ShikarjiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shikarji",
		TitleName: "尸佉罗时",
		TitleCode: "b_shikarji",
	}
}
