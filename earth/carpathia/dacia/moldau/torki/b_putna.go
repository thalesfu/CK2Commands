package torki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普特纳PutnaBarony struct {
	feud.BaseBarony
}

var BPutna普特纳 feud.Barony = &普特纳PutnaBarony{}

func init() {
    f := BPutna普特纳.(*普特纳PutnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "putna",
		TitleName: "普特纳",
		TitleCode: "b_putna",
	}
}
