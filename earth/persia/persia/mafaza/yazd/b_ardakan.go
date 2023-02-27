package yazd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔达坎ArdakanBarony struct {
	feud.BaseBarony
}

var BArdakan阿尔达坎 feud.Barony = &阿尔达坎ArdakanBarony{}

func init() {
    f := BArdakan阿尔达坎.(*阿尔达坎ArdakanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ardakan",
		TitleName: "阿尔达坎",
		TitleCode: "b_ardakan",
	}
}
