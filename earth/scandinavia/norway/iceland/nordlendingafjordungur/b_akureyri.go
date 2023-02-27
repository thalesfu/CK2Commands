package nordlendingafjordungur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克雷里AkureyriBarony struct {
	feud.BaseBarony
}

var BAkureyri阿克雷里 feud.Barony = &阿克雷里AkureyriBarony{}

func init() {
    f := BAkureyri阿克雷里.(*阿克雷里AkureyriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akureyri",
		TitleName: "阿克雷里",
		TitleCode: "b_akureyri",
	}
}
