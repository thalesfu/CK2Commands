package kyzyl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌拉甘UlaganBarony struct {
	feud.BaseBarony
}

var BUlagan乌拉甘 feud.Barony = &乌拉甘UlaganBarony{}

func init() {
    f := BUlagan乌拉甘.(*乌拉甘UlaganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulagan",
		TitleName: "乌拉甘",
		TitleCode: "b_ulagan",
	}
}
