package opole

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯切尔采StrzelceBarony struct {
	feud.BaseBarony
}

var BStrzelce斯切尔采 feud.Barony = &斯切尔采StrzelceBarony{}

func init() {
    f := BStrzelce斯切尔采.(*斯切尔采StrzelceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "strzelce",
		TitleName: "斯切尔采",
		TitleCode: "b_strzelce",
	}
}
