package craiova

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科拉比亚CorabiaBarony struct {
	feud.BaseBarony
}

var BCorabia科拉比亚 feud.Barony = &科拉比亚CorabiaBarony{}

func init() {
    f := BCorabia科拉比亚.(*科拉比亚CorabiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "corabia",
		TitleName: "科拉比亚",
		TitleCode: "b_corabia",
	}
}
