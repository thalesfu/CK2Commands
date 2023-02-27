package campulung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博罗瓦BovoraBarony struct {
	feud.BaseBarony
}

var BBovora博罗瓦 feud.Barony = &博罗瓦BovoraBarony{}

func init() {
    f := BBovora博罗瓦.(*博罗瓦BovoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bovora",
		TitleName: "博罗瓦",
		TitleCode: "b_bovora",
	}
}
