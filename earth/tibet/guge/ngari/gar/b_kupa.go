package gar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 过巴KupaBarony struct {
	feud.BaseBarony
}

var BKupa过巴 feud.Barony = &过巴KupaBarony{}

func init() {
	f := BKupa过巴.(*过巴KupaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kupa",
		TitleName: "过巴",
		TitleCode: "b_kupa",
	}
}
