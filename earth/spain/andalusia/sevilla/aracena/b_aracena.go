package aracena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉塞纳AracenaBarony struct {
	feud.BaseBarony
}

var BAracena阿拉塞纳 feud.Barony = &阿拉塞纳AracenaBarony{}

func init() {
    f := BAracena阿拉塞纳.(*阿拉塞纳AracenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aracena",
		TitleName: "阿拉塞纳",
		TitleCode: "b_aracena",
	}
}
