package ilam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉阿尔赫兰GhalehghiranBarony struct {
	feud.BaseBarony
}

var BGhalehghiran吉阿尔赫兰 feud.Barony = &吉阿尔赫兰GhalehghiranBarony{}

func init() {
    f := BGhalehghiran吉阿尔赫兰.(*吉阿尔赫兰GhalehghiranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghalehghiran",
		TitleName: "吉阿尔赫兰",
		TitleCode: "b_ghalehghiran",
	}
}
