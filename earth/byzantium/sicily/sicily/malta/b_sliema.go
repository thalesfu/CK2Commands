package malta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉巴特SliemaBarony struct {
	feud.BaseBarony
}

var BSliema拉巴特 feud.Barony = &拉巴特SliemaBarony{}

func init() {
    f := BSliema拉巴特.(*拉巴特SliemaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sliema",
		TitleName: "拉巴特",
		TitleCode: "b_sliema",
	}
}
