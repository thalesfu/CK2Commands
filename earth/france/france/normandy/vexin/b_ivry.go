package vexin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊夫里IvryBarony struct {
	feud.BaseBarony
}

var BIvry伊夫里 feud.Barony = &伊夫里IvryBarony{}

func init() {
    f := BIvry伊夫里.(*伊夫里IvryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ivry",
		TitleName: "伊夫里",
		TitleCode: "b_ivry",
	}
}
