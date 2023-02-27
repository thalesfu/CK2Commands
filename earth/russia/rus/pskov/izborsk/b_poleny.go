package izborsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波列内PolenyBarony struct {
	feud.BaseBarony
}

var BPoleny波列内 feud.Barony = &波列内PolenyBarony{}

func init() {
    f := BPoleny波列内.(*波列内PolenyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "poleny",
		TitleName: "波列内",
		TitleCode: "b_poleny",
	}
}
