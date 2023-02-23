package clydesdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伦弗鲁RenfrewBarony struct {
	feud.BaseBarony
}

var BRenfrew伦弗鲁 feud.Barony = &伦弗鲁RenfrewBarony{}

func init() {
	f := BRenfrew伦弗鲁.(*伦弗鲁RenfrewBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "renfrew",
		TitleName: "伦弗鲁",
		TitleCode: "b_renfrew",
	}
}
