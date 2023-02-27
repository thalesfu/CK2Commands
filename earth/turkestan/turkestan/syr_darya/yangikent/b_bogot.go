package yangikent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴哥特BogotBarony struct {
	feud.BaseBarony
}

var BBogot巴哥特 feud.Barony = &巴哥特BogotBarony{}

func init() {
    f := BBogot巴哥特.(*巴哥特BogotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bogot",
		TitleName: "巴哥特",
		TitleCode: "b_bogot",
	}
}
