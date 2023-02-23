package ostergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 灵斯塔霍尔姆RingstaholmBarony struct {
	feud.BaseBarony
}

var BRingstaholm灵斯塔霍尔姆 feud.Barony = &灵斯塔霍尔姆RingstaholmBarony{}

func init() {
	f := BRingstaholm灵斯塔霍尔姆.(*灵斯塔霍尔姆RingstaholmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ringstaholm",
		TitleName: "灵斯塔霍尔姆",
		TitleCode: "b_ringstaholm",
	}
}
