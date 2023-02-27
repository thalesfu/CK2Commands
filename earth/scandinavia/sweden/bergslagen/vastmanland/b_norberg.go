package vastmanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努尔贝里NorbergBarony struct {
	feud.BaseBarony
}

var BNorberg努尔贝里 feud.Barony = &努尔贝里NorbergBarony{}

func init() {
    f := BNorberg努尔贝里.(*努尔贝里NorbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "norberg",
		TitleName: "努尔贝里",
		TitleCode: "b_norberg",
	}
}
