package diskit

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托易斯ThoiseBarony struct {
	feud.BaseBarony
}

var BThoise托易斯 feud.Barony = &托易斯ThoiseBarony{}

func init() {
    f := BThoise托易斯.(*托易斯ThoiseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thoise",
		TitleName: "托易斯",
		TitleCode: "b_thoise",
	}
}
