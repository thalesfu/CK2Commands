package sambalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗尼补罗RanipurBarony struct {
	feud.BaseBarony
}

var BRanipur罗尼补罗 feud.Barony = &罗尼补罗RanipurBarony{}

func init() {
    f := BRanipur罗尼补罗.(*罗尼补罗RanipurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ranipur",
		TitleName: "罗尼补罗",
		TitleCode: "b_ranipur",
	}
}
