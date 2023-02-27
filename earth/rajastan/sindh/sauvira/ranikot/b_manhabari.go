package ranikot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼诃婆尼ManhabariBarony struct {
	feud.BaseBarony
}

var BManhabari曼诃婆尼 feud.Barony = &曼诃婆尼ManhabariBarony{}

func init() {
    f := BManhabari曼诃婆尼.(*曼诃婆尼ManhabariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manhabari",
		TitleName: "曼诃婆尼",
		TitleCode: "b_manhabari",
	}
}
