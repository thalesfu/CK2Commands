package rugen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普特布斯PutbusBarony struct {
	feud.BaseBarony
}

var BPutbus普特布斯 feud.Barony = &普特布斯PutbusBarony{}

func init() {
    f := BPutbus普特布斯.(*普特布斯PutbusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "putbus",
		TitleName: "普特布斯",
		TitleCode: "b_putbus",
	}
}
