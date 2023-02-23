package karluk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨雷沙甘SaryshaghanBarony struct {
	feud.BaseBarony
}

var BSaryshaghan萨雷沙甘 feud.Barony = &萨雷沙甘SaryshaghanBarony{}

func init() {
	f := BSaryshaghan萨雷沙甘.(*萨雷沙甘SaryshaghanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saryshaghan",
		TitleName: "萨雷沙甘",
		TitleCode: "b_saryshaghan",
	}
}
