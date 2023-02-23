package charsianon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈耳西亚农CharsianonBarony struct {
	feud.BaseBarony
}

var BCharsianon哈耳西亚农 feud.Barony = &哈耳西亚农CharsianonBarony{}

func init() {
	f := BCharsianon哈耳西亚农.(*哈耳西亚农CharsianonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charsianon",
		TitleName: "哈耳西亚农",
		TitleCode: "b_charsianon",
	}
}
