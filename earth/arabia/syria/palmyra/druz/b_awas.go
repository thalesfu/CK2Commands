package druz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿巴塞AwasBarony struct {
	feud.BaseBarony
}

var BAwas阿巴塞 feud.Barony = &阿巴塞AwasBarony{}

func init() {
    f := BAwas阿巴塞.(*阿巴塞AwasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "awas",
		TitleName: "阿巴塞",
		TitleCode: "b_awas",
	}
}
