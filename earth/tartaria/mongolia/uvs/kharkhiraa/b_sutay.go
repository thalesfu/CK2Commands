package kharkhiraa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏泰SutayBarony struct {
	feud.BaseBarony
}

var BSutay苏泰 feud.Barony = &苏泰SutayBarony{}

func init() {
    f := BSutay苏泰.(*苏泰SutayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sutay",
		TitleName: "苏泰",
		TitleCode: "b_sutay",
	}
}
