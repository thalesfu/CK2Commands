package wielunska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯切舒夫OstrzeszowBarony struct {
	feud.BaseBarony
}

var BOstrzeszow奥斯切舒夫 feud.Barony = &奥斯切舒夫OstrzeszowBarony{}

func init() {
    f := BOstrzeszow奥斯切舒夫.(*奥斯切舒夫OstrzeszowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ostrzeszow",
		TitleName: "奥斯切舒夫",
		TitleCode: "b_ostrzeszow",
	}
}
