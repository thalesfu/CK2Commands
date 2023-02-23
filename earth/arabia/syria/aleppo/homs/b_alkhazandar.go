package homs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈赞达AlkhazandarBarony struct {
	feud.BaseBarony
}

var BAlkhazandar哈赞达 feud.Barony = &哈赞达AlkhazandarBarony{}

func init() {
	f := BAlkhazandar哈赞达.(*哈赞达AlkhazandarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alkhazandar",
		TitleName: "哈赞达",
		TitleCode: "b_alkhazandar",
	}
}
