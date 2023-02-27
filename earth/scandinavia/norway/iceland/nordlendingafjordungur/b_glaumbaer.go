package nordlendingafjordungur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格伦拜尔GlaumbaerBarony struct {
	feud.BaseBarony
}

var BGlaumbaer格伦拜尔 feud.Barony = &格伦拜尔GlaumbaerBarony{}

func init() {
    f := BGlaumbaer格伦拜尔.(*格伦拜尔GlaumbaerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "glaumbaer",
		TitleName: "格伦拜尔",
		TitleCode: "b_glaumbaer",
	}
}
