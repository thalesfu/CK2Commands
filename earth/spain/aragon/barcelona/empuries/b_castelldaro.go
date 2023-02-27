package empuries

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿罗堡CastelldaroBarony struct {
	feud.BaseBarony
}

var BCastelldaro阿罗堡 feud.Barony = &阿罗堡CastelldaroBarony{}

func init() {
    f := BCastelldaro阿罗堡.(*阿罗堡CastelldaroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castelldaro",
		TitleName: "阿罗堡",
		TitleCode: "b_castelldaro",
	}
}
