package lubelska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 下卡齐米日NadwislaBarony struct {
	feud.BaseBarony
}

var BNadwisla下卡齐米日 feud.Barony = &下卡齐米日NadwislaBarony{}

func init() {
    f := BNadwisla下卡齐米日.(*下卡齐米日NadwislaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nadwisla",
		TitleName: "下卡齐米日",
		TitleCode: "b_nadwisla",
	}
}
