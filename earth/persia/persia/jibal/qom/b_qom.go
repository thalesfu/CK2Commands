package qom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库姆QomBarony struct {
	feud.BaseBarony
}

var BQom库姆 feud.Barony = &库姆QomBarony{}

func init() {
	f := BQom库姆.(*库姆QomBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qom",
		TitleName: "库姆",
		TitleCode: "b_qom",
	}
}
