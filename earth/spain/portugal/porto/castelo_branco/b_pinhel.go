package castelo_branco

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮涅尔PinhelBarony struct {
	feud.BaseBarony
}

var BPinhel皮涅尔 feud.Barony = &皮涅尔PinhelBarony{}

func init() {
    f := BPinhel皮涅尔.(*皮涅尔PinhelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pinhel",
		TitleName: "皮涅尔",
		TitleCode: "b_pinhel",
	}
}
