package bratslav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克雷若皮尔KryzhopilBarony struct {
	feud.BaseBarony
}

var BKryzhopil克雷若皮尔 feud.Barony = &克雷若皮尔KryzhopilBarony{}

func init() {
    f := BKryzhopil克雷若皮尔.(*克雷若皮尔KryzhopilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kryzhopil",
		TitleName: "克雷若皮尔",
		TitleCode: "b_kryzhopil",
	}
}
