package jiuquan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 肃州SuzhouBarony struct {
	feud.BaseBarony
}

var BSuzhou肃州 feud.Barony = &肃州SuzhouBarony{}

func init() {
    f := BSuzhou肃州.(*肃州SuzhouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suzhou",
		TitleName: "肃州",
		TitleCode: "b_suzhou",
	}
}
