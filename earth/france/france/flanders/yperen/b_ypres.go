package yperen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊普尔YpresBarony struct {
	feud.BaseBarony
}

var BYpres伊普尔 feud.Barony = &伊普尔YpresBarony{}

func init() {
    f := BYpres伊普尔.(*伊普尔YpresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ypres",
		TitleName: "伊普尔",
		TitleCode: "b_ypres",
	}
}
