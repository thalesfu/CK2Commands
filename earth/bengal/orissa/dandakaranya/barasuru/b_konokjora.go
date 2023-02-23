package barasuru

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘诺迦殊罗KonokjoraBarony struct {
	feud.BaseBarony
}

var BKonokjora拘诺迦殊罗 feud.Barony = &拘诺迦殊罗KonokjoraBarony{}

func init() {
	f := BKonokjora拘诺迦殊罗.(*拘诺迦殊罗KonokjoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "konokjora",
		TitleName: "拘诺迦殊罗",
		TitleCode: "b_konokjora",
	}
}
