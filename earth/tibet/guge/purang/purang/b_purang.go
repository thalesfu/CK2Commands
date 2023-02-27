package purang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布让PurangBarony struct {
	feud.BaseBarony
}

var BPurang布让 feud.Barony = &布让PurangBarony{}

func init() {
    f := BPurang布让.(*布让PurangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "purang",
		TitleName: "布让",
		TitleCode: "b_purang",
	}
}
