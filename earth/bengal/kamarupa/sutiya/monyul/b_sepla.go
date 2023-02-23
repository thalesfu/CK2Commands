package monyul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞巴SeplaBarony struct {
	feud.BaseBarony
}

var BSepla塞巴 feud.Barony = &塞巴SeplaBarony{}

func init() {
	f := BSepla塞巴.(*塞巴SeplaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sepla",
		TitleName: "塞巴",
		TitleCode: "b_sepla",
	}
}
