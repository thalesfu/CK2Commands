package badajoz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨夫拉ZafraBarony struct {
	feud.BaseBarony
}

var BZafra萨夫拉 feud.Barony = &萨夫拉ZafraBarony{}

func init() {
	f := BZafra萨夫拉.(*萨夫拉ZafraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zafra",
		TitleName: "萨夫拉",
		TitleCode: "b_zafra",
	}
}
