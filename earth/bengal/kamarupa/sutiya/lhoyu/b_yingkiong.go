package lhoyu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因基扬YingkiongBarony struct {
	feud.BaseBarony
}

var BYingkiong因基扬 feud.Barony = &因基扬YingkiongBarony{}

func init() {
	f := BYingkiong因基扬.(*因基扬YingkiongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yingkiong",
		TitleName: "因基扬",
		TitleCode: "b_yingkiong",
	}
}
