package bolshaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博利沙亚BolshayaBarony struct {
	feud.BaseBarony
}

var BBolshaya博利沙亚 feud.Barony = &博利沙亚BolshayaBarony{}

func init() {
    f := BBolshaya博利沙亚.(*博利沙亚BolshayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bolshaya",
		TitleName: "博利沙亚",
		TitleCode: "b_bolshaya",
	}
}
