package oshrusana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥什鲁萨纳OshrusanaBarony struct {
	feud.BaseBarony
}

var BOshrusana奥什鲁萨纳 feud.Barony = &奥什鲁萨纳OshrusanaBarony{}

func init() {
    f := BOshrusana奥什鲁萨纳.(*奥什鲁萨纳OshrusanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oshrusana",
		TitleName: "奥什鲁萨纳",
		TitleCode: "b_oshrusana",
	}
}
