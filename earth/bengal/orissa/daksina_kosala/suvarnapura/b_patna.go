package suvarnapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波吒那PatnaBarony struct {
	feud.BaseBarony
}

var BPatna波吒那 feud.Barony = &波吒那PatnaBarony{}

func init() {
    f := BPatna波吒那.(*波吒那PatnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "patna",
		TitleName: "波吒那",
		TitleCode: "b_patna",
	}
}
