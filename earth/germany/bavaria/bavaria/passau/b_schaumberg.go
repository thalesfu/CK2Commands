package passau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 绍姆堡SchaumbergBarony struct {
	feud.BaseBarony
}

var BSchaumberg绍姆堡 feud.Barony = &绍姆堡SchaumbergBarony{}

func init() {
    f := BSchaumberg绍姆堡.(*绍姆堡SchaumbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schaumberg",
		TitleName: "绍姆堡",
		TitleCode: "b_schaumberg",
	}
}
