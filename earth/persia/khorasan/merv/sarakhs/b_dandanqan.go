package sarakhs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丹丹坎DandanqanBarony struct {
	feud.BaseBarony
}

var BDandanqan丹丹坎 feud.Barony = &丹丹坎DandanqanBarony{}

func init() {
    f := BDandanqan丹丹坎.(*丹丹坎DandanqanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dandanqan",
		TitleName: "丹丹坎",
		TitleCode: "b_dandanqan",
	}
}
