package kermanshah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库泽兰KuzaranBarony struct {
	feud.BaseBarony
}

var BKuzaran库泽兰 feud.Barony = &库泽兰KuzaranBarony{}

func init() {
	f := BKuzaran库泽兰.(*库泽兰KuzaranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuzaran",
		TitleName: "库泽兰",
		TitleCode: "b_kuzaran",
	}
}
