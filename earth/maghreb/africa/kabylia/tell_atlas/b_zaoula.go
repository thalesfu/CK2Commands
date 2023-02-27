package tell_atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祖拉ZaoulaBarony struct {
	feud.BaseBarony
}

var BZaoula祖拉 feud.Barony = &祖拉ZaoulaBarony{}

func init() {
    f := BZaoula祖拉.(*祖拉ZaoulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaoula",
		TitleName: "祖拉",
		TitleCode: "b_zaoula",
	}
}
