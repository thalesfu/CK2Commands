package minsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅斯维日NyasvizhBarony struct {
	feud.BaseBarony
}

var BNyasvizh涅斯维日 feud.Barony = &涅斯维日NyasvizhBarony{}

func init() {
    f := BNyasvizh涅斯维日.(*涅斯维日NyasvizhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyasvizh",
		TitleName: "涅斯维日",
		TitleCode: "b_nyasvizh",
	}
}
