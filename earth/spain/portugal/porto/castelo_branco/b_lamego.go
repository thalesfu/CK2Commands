package castelo_branco

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉梅戈LamegoBarony struct {
	feud.BaseBarony
}

var BLamego拉梅戈 feud.Barony = &拉梅戈LamegoBarony{}

func init() {
    f := BLamego拉梅戈.(*拉梅戈LamegoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lamego",
		TitleName: "拉梅戈",
		TitleCode: "b_lamego",
	}
}
