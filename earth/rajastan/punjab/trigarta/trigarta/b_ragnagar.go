package trigarta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗阇那揭罗RagnagarBarony struct {
	feud.BaseBarony
}

var BRagnagar罗阇那揭罗 feud.Barony = &罗阇那揭罗RagnagarBarony{}

func init() {
    f := BRagnagar罗阇那揭罗.(*罗阇那揭罗RagnagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ragnagar",
		TitleName: "罗阇那揭罗",
		TitleCode: "b_ragnagar",
	}
}
