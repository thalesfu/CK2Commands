package borovichi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博罗维奇BorovichiBarony struct {
	feud.BaseBarony
}

var BBorovichi博罗维奇 feud.Barony = &博罗维奇BorovichiBarony{}

func init() {
    f := BBorovichi博罗维奇.(*博罗维奇BorovichiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borovichi",
		TitleName: "博罗维奇",
		TitleCode: "b_borovichi",
	}
}
