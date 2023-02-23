package middlesex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈罗HarrowBarony struct {
	feud.BaseBarony
}

var BHarrow哈罗 feud.Barony = &哈罗HarrowBarony{}

func init() {
	f := BHarrow哈罗.(*哈罗HarrowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harrow",
		TitleName: "哈罗",
		TitleCode: "b_harrow",
	}
}
