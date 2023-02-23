package geneve

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊瓦尔YvoireBarony struct {
	feud.BaseBarony
}

var BYvoire伊瓦尔 feud.Barony = &伊瓦尔YvoireBarony{}

func init() {
	f := BYvoire伊瓦尔.(*伊瓦尔YvoireBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yvoire",
		TitleName: "伊瓦尔",
		TitleCode: "b_yvoire",
	}
}
