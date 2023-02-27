package cholamandalam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 底哩尸罗钵利TiruccirapalliBarony struct {
	feud.BaseBarony
}

var BTiruccirapalli底哩尸罗钵利 feud.Barony = &底哩尸罗钵利TiruccirapalliBarony{}

func init() {
    f := BTiruccirapalli底哩尸罗钵利.(*底哩尸罗钵利TiruccirapalliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiruccirapalli",
		TitleName: "底哩尸罗钵利",
		TitleCode: "b_tiruccirapalli",
	}
}
