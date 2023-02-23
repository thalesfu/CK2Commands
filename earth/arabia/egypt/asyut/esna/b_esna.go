package esna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯纳EsnaBarony struct {
	feud.BaseBarony
}

var BEsna埃斯纳 feud.Barony = &埃斯纳EsnaBarony{}

func init() {
	f := BEsna埃斯纳.(*埃斯纳EsnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "esna",
		TitleName: "埃斯纳",
		TitleCode: "b_esna",
	}
}
