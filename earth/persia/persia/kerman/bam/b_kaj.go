package bam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡吉KajBarony struct {
	feud.BaseBarony
}

var BKaj卡吉 feud.Barony = &卡吉KajBarony{}

func init() {
    f := BKaj卡吉.(*卡吉KajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaj",
		TitleName: "卡吉",
		TitleCode: "b_kaj",
	}
}
