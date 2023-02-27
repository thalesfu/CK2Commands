package kuma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库里_尤尔特KuliyurtBarony struct {
	feud.BaseBarony
}

var BKuliyurt库里_尤尔特 feud.Barony = &库里_尤尔特KuliyurtBarony{}

func init() {
    f := BKuliyurt库里_尤尔特.(*库里_尤尔特KuliyurtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuliyurt",
		TitleName: "库里_尤尔特",
		TitleCode: "b_kuliyurt",
	}
}
