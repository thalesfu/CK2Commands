package karluk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弩失毕NushibiBarony struct {
	feud.BaseBarony
}

var BNushibi弩失毕 feud.Barony = &弩失毕NushibiBarony{}

func init() {
	f := BNushibi弩失毕.(*弩失毕NushibiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nushibi",
		TitleName: "弩失毕",
		TitleCode: "b_nushibi",
	}
}
