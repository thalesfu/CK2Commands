package rayapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗卢豆RablutoBarony struct {
	feud.BaseBarony
}

var BRabluto罗卢豆 feud.Barony = &罗卢豆RablutoBarony{}

func init() {
    f := BRabluto罗卢豆.(*罗卢豆RablutoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rabluto",
		TitleName: "罗卢豆",
		TitleCode: "b_rabluto",
	}
}
