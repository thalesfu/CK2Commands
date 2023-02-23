package hajr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜勒朱尔希BaljurashiBarony struct {
	feud.BaseBarony
}

var BBaljurashi拜勒朱尔希 feud.Barony = &拜勒朱尔希BaljurashiBarony{}

func init() {
	f := BBaljurashi拜勒朱尔希.(*拜勒朱尔希BaljurashiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baljurashi",
		TitleName: "拜勒朱尔希",
		TitleCode: "b_baljurashi",
	}
}
