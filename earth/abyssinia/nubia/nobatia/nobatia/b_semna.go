package nobatia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞姆纳SemnaBarony struct {
	feud.BaseBarony
}

var BSemna塞姆纳 feud.Barony = &塞姆纳SemnaBarony{}

func init() {
    f := BSemna塞姆纳.(*塞姆纳SemnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "semna",
		TitleName: "塞姆纳",
		TitleCode: "b_semna",
	}
}
