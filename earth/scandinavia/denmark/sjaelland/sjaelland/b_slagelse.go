package sjaelland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯劳厄尔瑟SlagelseBarony struct {
	feud.BaseBarony
}

var BSlagelse斯劳厄尔瑟 feud.Barony = &斯劳厄尔瑟SlagelseBarony{}

func init() {
    f := BSlagelse斯劳厄尔瑟.(*斯劳厄尔瑟SlagelseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "slagelse",
		TitleName: "斯劳厄尔瑟",
		TitleCode: "b_slagelse",
	}
}
