package uchturpan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 东凌井TunglingchingBarony struct {
	feud.BaseBarony
}

var BTunglingching东凌井 feud.Barony = &东凌井TunglingchingBarony{}

func init() {
	f := BTunglingching东凌井.(*东凌井TunglingchingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tunglingching",
		TitleName: "东凌井",
		TitleCode: "b_tunglingching",
	}
}
