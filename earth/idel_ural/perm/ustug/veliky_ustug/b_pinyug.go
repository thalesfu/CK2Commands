package veliky_ustug

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮纽格PinyugBarony struct {
	feud.BaseBarony
}

var BPinyug皮纽格 feud.Barony = &皮纽格PinyugBarony{}

func init() {
    f := BPinyug皮纽格.(*皮纽格PinyugBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pinyug",
		TitleName: "皮纽格",
		TitleCode: "b_pinyug",
	}
}
