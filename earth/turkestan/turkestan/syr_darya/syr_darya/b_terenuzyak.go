package syr_darya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 捷连_乌兹亚克TerenuzyakBarony struct {
	feud.BaseBarony
}

var BTerenuzyak捷连_乌兹亚克 feud.Barony = &捷连_乌兹亚克TerenuzyakBarony{}

func init() {
    f := BTerenuzyak捷连_乌兹亚克.(*捷连_乌兹亚克TerenuzyakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "terenuzyak",
		TitleName: "捷连_乌兹亚克",
		TitleCode: "b_terenuzyak",
	}
}
