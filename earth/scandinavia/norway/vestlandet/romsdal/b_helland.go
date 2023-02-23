package romsdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫兰HellandBarony struct {
	feud.BaseBarony
}

var BHelland赫兰 feud.Barony = &赫兰HellandBarony{}

func init() {
	f := BHelland赫兰.(*赫兰HellandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "helland",
		TitleName: "赫兰",
		TitleCode: "b_helland",
	}
}
