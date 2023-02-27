package lhatok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乍丫ZhagyabBarony struct {
	feud.BaseBarony
}

var BZhagyab乍丫 feud.Barony = &乍丫ZhagyabBarony{}

func init() {
    f := BZhagyab乍丫.(*乍丫ZhagyabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhagyab",
		TitleName: "乍丫",
		TitleCode: "b_zhagyab",
	}
}
