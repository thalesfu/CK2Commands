package escuens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗龙特奈FrontenayBarony struct {
	feud.BaseBarony
}

var BFrontenay弗龙特奈 feud.Barony = &弗龙特奈FrontenayBarony{}

func init() {
    f := BFrontenay弗龙特奈.(*弗龙特奈FrontenayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "frontenay",
		TitleName: "弗龙特奈",
		TitleCode: "b_frontenay",
	}
}
