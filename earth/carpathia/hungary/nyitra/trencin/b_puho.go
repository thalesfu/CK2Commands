package trencin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普霍PuhoBarony struct {
	feud.BaseBarony
}

var BPuho普霍 feud.Barony = &普霍PuhoBarony{}

func init() {
	f := BPuho普霍.(*普霍PuhoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "puho",
		TitleName: "普霍",
		TitleCode: "b_puho",
	}
}
