package osterreich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴登Baden_wienBarony struct {
	feud.BaseBarony
}

var BBaden_wien巴登 feud.Barony = &巴登Baden_wienBarony{}

func init() {
    f := BBaden_wien巴登.(*巴登Baden_wienBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baden_wien",
		TitleName: "巴登",
		TitleCode: "b_baden_wien",
	}
}
