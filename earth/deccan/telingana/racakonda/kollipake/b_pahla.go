package kollipake

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波诃罗PahlaBarony struct {
	feud.BaseBarony
}

var BPahla波诃罗 feud.Barony = &波诃罗PahlaBarony{}

func init() {
    f := BPahla波诃罗.(*波诃罗PahlaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pahla",
		TitleName: "波诃罗",
		TitleCode: "b_pahla",
	}
}
