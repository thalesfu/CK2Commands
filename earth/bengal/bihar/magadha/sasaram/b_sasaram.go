package sasaram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨娑蓝SasaramBarony struct {
	feud.BaseBarony
}

var BSasaram萨娑蓝 feud.Barony = &萨娑蓝SasaramBarony{}

func init() {
	f := BSasaram萨娑蓝.(*萨娑蓝SasaramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sasaram",
		TitleName: "萨娑蓝",
		TitleCode: "b_sasaram",
	}
}
