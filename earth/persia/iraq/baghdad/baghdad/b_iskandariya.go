package baghdad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯坎德里耶IskandariyaBarony struct {
	feud.BaseBarony
}

var BIskandariya伊斯坎德里耶 feud.Barony = &伊斯坎德里耶IskandariyaBarony{}

func init() {
	f := BIskandariya伊斯坎德里耶.(*伊斯坎德里耶IskandariyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iskandariya",
		TitleName: "伊斯坎德里耶",
		TitleCode: "b_iskandariya",
	}
}
