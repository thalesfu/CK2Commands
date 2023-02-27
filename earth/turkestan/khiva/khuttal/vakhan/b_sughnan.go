package vakhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 识匿SughnanBarony struct {
	feud.BaseBarony
}

var BSughnan识匿 feud.Barony = &识匿SughnanBarony{}

func init() {
    f := BSughnan识匿.(*识匿SughnanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sughnan",
		TitleName: "识匿",
		TitleCode: "b_sughnan",
	}
}
