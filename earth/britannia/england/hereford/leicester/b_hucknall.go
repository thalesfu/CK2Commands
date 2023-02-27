package leicester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈克纳尔HucknallBarony struct {
	feud.BaseBarony
}

var BHucknall哈克纳尔 feud.Barony = &哈克纳尔HucknallBarony{}

func init() {
    f := BHucknall哈克纳尔.(*哈克纳尔HucknallBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hucknall",
		TitleName: "哈克纳尔",
		TitleCode: "b_hucknall",
	}
}
