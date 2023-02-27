package draa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾卢古姆AlogoumBarony struct {
	feud.BaseBarony
}

var BAlogoum艾卢古姆 feud.Barony = &艾卢古姆AlogoumBarony{}

func init() {
    f := BAlogoum艾卢古姆.(*艾卢古姆AlogoumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alogoum",
		TitleName: "艾卢古姆",
		TitleCode: "b_alogoum",
	}
}
