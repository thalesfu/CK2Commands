package qohistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔达斯坎BardaskanBarony struct {
	feud.BaseBarony
}

var BBardaskan巴尔达斯坎 feud.Barony = &巴尔达斯坎BardaskanBarony{}

func init() {
    f := BBardaskan巴尔达斯坎.(*巴尔达斯坎BardaskanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bardaskan",
		TitleName: "巴尔达斯坎",
		TitleCode: "b_bardaskan",
	}
}
