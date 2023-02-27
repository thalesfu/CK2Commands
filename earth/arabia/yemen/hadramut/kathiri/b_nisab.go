package kathiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼萨卜NisabBarony struct {
	feud.BaseBarony
}

var BNisab尼萨卜 feud.Barony = &尼萨卜NisabBarony{}

func init() {
    f := BNisab尼萨卜.(*尼萨卜NisabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nisab",
		TitleName: "尼萨卜",
		TitleCode: "b_nisab",
	}
}
