package utva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨雷乌米尔SaryumirBarony struct {
	feud.BaseBarony
}

var BSaryumir萨雷乌米尔 feud.Barony = &萨雷乌米尔SaryumirBarony{}

func init() {
    f := BSaryumir萨雷乌米尔.(*萨雷乌米尔SaryumirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saryumir",
		TitleName: "萨雷乌米尔",
		TitleCode: "b_saryumir",
	}
}
