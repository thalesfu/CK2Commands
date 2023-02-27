package oberbayern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普法芬贝格PfaffenbergBarony struct {
	feud.BaseBarony
}

var BPfaffenberg普法芬贝格 feud.Barony = &普法芬贝格PfaffenbergBarony{}

func init() {
    f := BPfaffenberg普法芬贝格.(*普法芬贝格PfaffenbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pfaffenberg",
		TitleName: "普法芬贝格",
		TitleCode: "b_pfaffenberg",
	}
}
