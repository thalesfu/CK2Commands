package kara_bogaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡占KadzhanBarony struct {
	feud.BaseBarony
}

var BKadzhan卡占 feud.Barony = &卡占KadzhanBarony{}

func init() {
    f := BKadzhan卡占.(*卡占KadzhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kadzhan",
		TitleName: "卡占",
		TitleCode: "b_kadzhan",
	}
}
