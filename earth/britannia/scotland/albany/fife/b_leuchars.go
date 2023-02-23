package fife

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢赫斯LeucharsBarony struct {
	feud.BaseBarony
}

var BLeuchars卢赫斯 feud.Barony = &卢赫斯LeucharsBarony{}

func init() {
	f := BLeuchars卢赫斯.(*卢赫斯LeucharsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leuchars",
		TitleName: "卢赫斯",
		TitleCode: "b_leuchars",
	}
}
