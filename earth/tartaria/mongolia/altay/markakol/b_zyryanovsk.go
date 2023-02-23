package markakol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 济良诺夫斯克ZyryanovskBarony struct {
	feud.BaseBarony
}

var BZyryanovsk济良诺夫斯克 feud.Barony = &济良诺夫斯克ZyryanovskBarony{}

func init() {
	f := BZyryanovsk济良诺夫斯克.(*济良诺夫斯克ZyryanovskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zyryanovsk",
		TitleName: "济良诺夫斯克",
		TitleCode: "b_zyryanovsk",
	}
}
