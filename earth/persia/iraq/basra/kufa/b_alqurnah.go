package kufa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尔奈AlqurnahBarony struct {
	feud.BaseBarony
}

var BAlqurnah古尔奈 feud.Barony = &古尔奈AlqurnahBarony{}

func init() {
	f := BAlqurnah古尔奈.(*古尔奈AlqurnahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alqurnah",
		TitleName: "古尔奈",
		TitleCode: "b_alqurnah",
	}
}
