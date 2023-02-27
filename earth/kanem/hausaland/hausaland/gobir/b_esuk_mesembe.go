package gobir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃苏克梅森贝Esuk_mesembeBarony struct {
	feud.BaseBarony
}

var BEsuk_mesembe埃苏克梅森贝 feud.Barony = &埃苏克梅森贝Esuk_mesembeBarony{}

func init() {
    f := BEsuk_mesembe埃苏克梅森贝.(*埃苏克梅森贝Esuk_mesembeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "esuk_mesembe",
		TitleName: "埃苏克梅森贝",
		TitleCode: "b_esuk_mesembe",
	}
}
