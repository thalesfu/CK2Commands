package candradvipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波迦罗犍阇BakarganjBarony struct {
	feud.BaseBarony
}

var BBakarganj波迦罗犍阇 feud.Barony = &波迦罗犍阇BakarganjBarony{}

func init() {
	f := BBakarganj波迦罗犍阇.(*波迦罗犍阇BakarganjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bakarganj",
		TitleName: "波迦罗犍阇",
		TitleCode: "b_bakarganj",
	}
}
