package leoben

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布鲁克Bruck_leobenBarony struct {
	feud.BaseBarony
}

var BBruck_leoben布鲁克 feud.Barony = &布鲁克Bruck_leobenBarony{}

func init() {
    f := BBruck_leoben布鲁克.(*布鲁克Bruck_leobenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bruck_leoben",
		TitleName: "布鲁克",
		TitleCode: "b_bruck_leoben",
	}
}
