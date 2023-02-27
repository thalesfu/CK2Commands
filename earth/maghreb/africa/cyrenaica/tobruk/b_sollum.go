package tobruk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞卢姆SollumBarony struct {
	feud.BaseBarony
}

var BSollum塞卢姆 feud.Barony = &塞卢姆SollumBarony{}

func init() {
    f := BSollum塞卢姆.(*塞卢姆SollumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sollum",
		TitleName: "塞卢姆",
		TitleCode: "b_sollum",
	}
}
