package paro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 祈楚KyichuBarony struct {
	feud.BaseBarony
}

var BKyichu祈楚 feud.Barony = &祈楚KyichuBarony{}

func init() {
    f := BKyichu祈楚.(*祈楚KyichuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyichu",
		TitleName: "祈楚",
		TitleCode: "b_kyichu",
	}
}
