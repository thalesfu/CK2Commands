package burhanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 补罗汉补罗BurhanpurBarony struct {
	feud.BaseBarony
}

var BBurhanpur补罗汉补罗 feud.Barony = &补罗汉补罗BurhanpurBarony{}

func init() {
	f := BBurhanpur补罗汉补罗.(*补罗汉补罗BurhanpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burhanpur",
		TitleName: "补罗汉补罗",
		TitleCode: "b_burhanpur",
	}
}
