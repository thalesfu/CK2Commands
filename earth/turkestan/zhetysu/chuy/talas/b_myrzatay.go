package talas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅尔扎泰MyrzatayBarony struct {
	feud.BaseBarony
}

var BMyrzatay梅尔扎泰 feud.Barony = &梅尔扎泰MyrzatayBarony{}

func init() {
    f := BMyrzatay梅尔扎泰.(*梅尔扎泰MyrzatayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "myrzatay",
		TitleName: "梅尔扎泰",
		TitleCode: "b_myrzatay",
	}
}
