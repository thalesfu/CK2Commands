package torzhok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 济济诺ZizinoBarony struct {
	feud.BaseBarony
}

var BZizino济济诺 feud.Barony = &济济诺ZizinoBarony{}

func init() {
	f := BZizino济济诺.(*济济诺ZizinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zizino",
		TitleName: "济济诺",
		TitleCode: "b_zizino",
	}
}
