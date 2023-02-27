package scalovia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉格尼特RagnitBarony struct {
	feud.BaseBarony
}

var BRagnit拉格尼特 feud.Barony = &拉格尼特RagnitBarony{}

func init() {
    f := BRagnit拉格尼特.(*拉格尼特RagnitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ragnit",
		TitleName: "拉格尼特",
		TitleCode: "b_ragnit",
	}
}
