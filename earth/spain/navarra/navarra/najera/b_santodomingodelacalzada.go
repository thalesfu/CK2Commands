package najera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣多明各德拉卡尔萨达SantodomingodelacalzadaBarony struct {
	feud.BaseBarony
}

var BSantodomingodelacalzada圣多明各德拉卡尔萨达 feud.Barony = &圣多明各德拉卡尔萨达SantodomingodelacalzadaBarony{}

func init() {
    f := BSantodomingodelacalzada圣多明各德拉卡尔萨达.(*圣多明各德拉卡尔萨达SantodomingodelacalzadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "santodomingodelacalzada",
		TitleName: "圣多明各德拉卡尔萨达",
		TitleCode: "b_santodomingodelacalzada",
	}
}
