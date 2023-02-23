package seletyteniz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨雷奥巴SaryobaBarony struct {
	feud.BaseBarony
}

var BSaryoba萨雷奥巴 feud.Barony = &萨雷奥巴SaryobaBarony{}

func init() {
	f := BSaryoba萨雷奥巴.(*萨雷奥巴SaryobaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saryoba",
		TitleName: "萨雷奥巴",
		TitleCode: "b_saryoba",
	}
}
