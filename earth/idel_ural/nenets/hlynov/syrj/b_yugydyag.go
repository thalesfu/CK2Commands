package syrj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤格德亚格YugydyagBarony struct {
	feud.BaseBarony
}

var BYugydyag尤格德亚格 feud.Barony = &尤格德亚格YugydyagBarony{}

func init() {
    f := BYugydyag尤格德亚格.(*尤格德亚格YugydyagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yugydyag",
		TitleName: "尤格德亚格",
		TitleCode: "b_yugydyag",
	}
}
