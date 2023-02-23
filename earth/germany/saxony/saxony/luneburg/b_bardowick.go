package luneburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴多维克BardowickBarony struct {
	feud.BaseBarony
}

var BBardowick巴多维克 feud.Barony = &巴多维克BardowickBarony{}

func init() {
	f := BBardowick巴多维克.(*巴多维克BardowickBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bardowick",
		TitleName: "巴多维克",
		TitleCode: "b_bardowick",
	}
}
