package gilgit

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波呼特PahotBarony struct {
	feud.BaseBarony
}

var BPahot波呼特 feud.Barony = &波呼特PahotBarony{}

func init() {
    f := BPahot波呼特.(*波呼特PahotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pahot",
		TitleName: "波呼特",
		TitleCode: "b_pahot",
	}
}
