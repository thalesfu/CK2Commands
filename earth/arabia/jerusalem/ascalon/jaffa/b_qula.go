package jaffa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔拉QulaBarony struct {
	feud.BaseBarony
}

var BQula库尔拉 feud.Barony = &库尔拉QulaBarony{}

func init() {
	f := BQula库尔拉.(*库尔拉QulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qula",
		TitleName: "库尔拉",
		TitleCode: "b_qula",
	}
}
