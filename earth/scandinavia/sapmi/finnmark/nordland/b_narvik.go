package nordland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳尔维克NarvikBarony struct {
	feud.BaseBarony
}

var BNarvik纳尔维克 feud.Barony = &纳尔维克NarvikBarony{}

func init() {
    f := BNarvik纳尔维克.(*纳尔维克NarvikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narvik",
		TitleName: "纳尔维克",
		TitleCode: "b_narvik",
	}
}
