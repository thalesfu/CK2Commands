package veglia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尔克KrkBarony struct {
	feud.BaseBarony
}

var BKrk克尔克 feud.Barony = &克尔克KrkBarony{}

func init() {
	f := BKrk克尔克.(*克尔克KrkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krk",
		TitleName: "克尔克",
		TitleCode: "b_krk",
	}
}
