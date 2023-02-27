package najran

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里巴哈RibakhahBarony struct {
	feud.BaseBarony
}

var BRibakhah里巴哈 feud.Barony = &里巴哈RibakhahBarony{}

func init() {
    f := BRibakhah里巴哈.(*里巴哈RibakhahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ribakhah",
		TitleName: "里巴哈",
		TitleCode: "b_ribakhah",
	}
}
