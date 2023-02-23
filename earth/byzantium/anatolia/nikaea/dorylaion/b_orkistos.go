package dorylaion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧基斯图斯OrkistosBarony struct {
	feud.BaseBarony
}

var BOrkistos欧基斯图斯 feud.Barony = &欧基斯图斯OrkistosBarony{}

func init() {
	f := BOrkistos欧基斯图斯.(*欧基斯图斯OrkistosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orkistos",
		TitleName: "欧基斯图斯",
		TitleCode: "b_orkistos",
	}
}
