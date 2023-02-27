package hastinapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因陀罗萨傥那IndrasthanaBarony struct {
	feud.BaseBarony
}

var BIndrasthana因陀罗萨傥那 feud.Barony = &因陀罗萨傥那IndrasthanaBarony{}

func init() {
    f := BIndrasthana因陀罗萨傥那.(*因陀罗萨傥那IndrasthanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "indrasthana",
		TitleName: "因陀罗萨傥那",
		TitleCode: "b_indrasthana",
	}
}
