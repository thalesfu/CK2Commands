package armail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普咎柯PejokeBarony struct {
	feud.BaseBarony
}

var BPejoke普咎柯 feud.Barony = &普咎柯PejokeBarony{}

func init() {
    f := BPejoke普咎柯.(*普咎柯PejokeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pejoke",
		TitleName: "普咎柯",
		TitleCode: "b_pejoke",
	}
}
