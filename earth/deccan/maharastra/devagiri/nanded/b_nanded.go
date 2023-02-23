package nanded

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 难提陀NandedBarony struct {
	feud.BaseBarony
}

var BNanded难提陀 feud.Barony = &难提陀NandedBarony{}

func init() {
	f := BNanded难提陀.(*难提陀NandedBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nanded",
		TitleName: "难提陀",
		TitleCode: "b_nanded",
	}
}
