package winchester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 惠韦尔WherwellBarony struct {
	feud.BaseBarony
}

var BWherwell惠韦尔 feud.Barony = &惠韦尔WherwellBarony{}

func init() {
	f := BWherwell惠韦尔.(*惠韦尔WherwellBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wherwell",
		TitleName: "惠韦尔",
		TitleCode: "b_wherwell",
	}
}
