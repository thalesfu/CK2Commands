package bergenshus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特兰维克StrandvikBarony struct {
	feud.BaseBarony
}

var BStrandvik斯特兰维克 feud.Barony = &斯特兰维克StrandvikBarony{}

func init() {
    f := BStrandvik斯特兰维克.(*斯特兰维克StrandvikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "strandvik",
		TitleName: "斯特兰维克",
		TitleCode: "b_strandvik",
	}
}
