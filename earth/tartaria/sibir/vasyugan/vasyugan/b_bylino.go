package vasyugan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝利诺BylinoBarony struct {
	feud.BaseBarony
}

var BBylino贝利诺 feud.Barony = &贝利诺BylinoBarony{}

func init() {
    f := BBylino贝利诺.(*贝利诺BylinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bylino",
		TitleName: "贝利诺",
		TitleCode: "b_bylino",
	}
}
