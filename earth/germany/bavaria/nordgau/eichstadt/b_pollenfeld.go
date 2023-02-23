package eichstadt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波伦费尔德PollenfeldBarony struct {
	feud.BaseBarony
}

var BPollenfeld波伦费尔德 feud.Barony = &波伦费尔德PollenfeldBarony{}

func init() {
	f := BPollenfeld波伦费尔德.(*波伦费尔德PollenfeldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pollenfeld",
		TitleName: "波伦费尔德",
		TitleCode: "b_pollenfeld",
	}
}
