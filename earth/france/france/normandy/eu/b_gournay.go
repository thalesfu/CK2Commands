package eu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尔奈GournayBarony struct {
	feud.BaseBarony
}

var BGournay古尔奈 feud.Barony = &古尔奈GournayBarony{}

func init() {
    f := BGournay古尔奈.(*古尔奈GournayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gournay",
		TitleName: "古尔奈",
		TitleCode: "b_gournay",
	}
}
