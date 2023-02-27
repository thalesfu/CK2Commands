package urzhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔札尔UrzharBarony struct {
	feud.BaseBarony
}

var BUrzhar乌尔札尔 feud.Barony = &乌尔札尔UrzharBarony{}

func init() {
    f := BUrzhar乌尔札尔.(*乌尔札尔UrzharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urzhar",
		TitleName: "乌尔札尔",
		TitleCode: "b_urzhar",
	}
}
