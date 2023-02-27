package iarmond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因尼斯福伦InnisfallenBarony struct {
	feud.BaseBarony
}

var BInnisfallen因尼斯福伦 feud.Barony = &因尼斯福伦InnisfallenBarony{}

func init() {
    f := BInnisfallen因尼斯福伦.(*因尼斯福伦InnisfallenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "innisfallen",
		TitleName: "因尼斯福伦",
		TitleCode: "b_innisfallen",
	}
}
