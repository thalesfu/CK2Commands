package nagormo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 新乐XinleBarony struct {
	feud.BaseBarony
}

var BXinle新乐 feud.Barony = &新乐XinleBarony{}

func init() {
    f := BXinle新乐.(*新乐XinleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xinle",
		TitleName: "新乐",
		TitleCode: "b_xinle",
	}
}
