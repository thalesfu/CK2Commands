package algeciras

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯特波纳EsteponaBarony struct {
	feud.BaseBarony
}

var BEstepona埃斯特波纳 feud.Barony = &埃斯特波纳EsteponaBarony{}

func init() {
    f := BEstepona埃斯特波纳.(*埃斯特波纳EsteponaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "estepona",
		TitleName: "埃斯特波纳",
		TitleCode: "b_estepona",
	}
}
