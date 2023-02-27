package samarkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟底痕IshtikhonBarony struct {
	feud.BaseBarony
}

var BIshtikhon瑟底痕 feud.Barony = &瑟底痕IshtikhonBarony{}

func init() {
    f := BIshtikhon瑟底痕.(*瑟底痕IshtikhonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ishtikhon",
		TitleName: "瑟底痕",
		TitleCode: "b_ishtikhon",
	}
}
