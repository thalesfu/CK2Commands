package karelen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆耶泽尔斯基MuezerskyBarony struct {
	feud.BaseBarony
}

var BMuezersky穆耶泽尔斯基 feud.Barony = &穆耶泽尔斯基MuezerskyBarony{}

func init() {
    f := BMuezersky穆耶泽尔斯基.(*穆耶泽尔斯基MuezerskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muezersky",
		TitleName: "穆耶泽尔斯基",
		TitleCode: "b_muezersky",
	}
}
