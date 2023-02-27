package kalyani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那罗衍那补罗NarayanapurBarony struct {
	feud.BaseBarony
}

var BNarayanapur那罗衍那补罗 feud.Barony = &那罗衍那补罗NarayanapurBarony{}

func init() {
    f := BNarayanapur那罗衍那补罗.(*那罗衍那补罗NarayanapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narayanapur",
		TitleName: "那罗衍那补罗",
		TitleCode: "b_narayanapur",
	}
}
