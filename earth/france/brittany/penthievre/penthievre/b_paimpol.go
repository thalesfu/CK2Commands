package penthievre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 潘波勒PaimpolBarony struct {
	feud.BaseBarony
}

var BPaimpol潘波勒 feud.Barony = &潘波勒PaimpolBarony{}

func init() {
    f := BPaimpol潘波勒.(*潘波勒PaimpolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paimpol",
		TitleName: "潘波勒",
		TitleCode: "b_paimpol",
	}
}
