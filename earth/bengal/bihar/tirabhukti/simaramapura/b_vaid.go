package simaramapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗陀VaidBarony struct {
	feud.BaseBarony
}

var BVaid毗陀 feud.Barony = &毗陀VaidBarony{}

func init() {
    f := BVaid毗陀.(*毗陀VaidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vaid",
		TitleName: "毗陀",
		TitleCode: "b_vaid",
	}
}
