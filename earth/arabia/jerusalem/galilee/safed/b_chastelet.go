package safed

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙特莱ChasteletBarony struct {
	feud.BaseBarony
}

var BChastelet沙特莱 feud.Barony = &沙特莱ChasteletBarony{}

func init() {
    f := BChastelet沙特莱.(*沙特莱ChasteletBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chastelet",
		TitleName: "沙特莱",
		TitleCode: "b_chastelet",
	}
}
