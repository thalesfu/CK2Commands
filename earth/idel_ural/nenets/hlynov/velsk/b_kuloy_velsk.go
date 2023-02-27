package velsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库洛伊Kuloy_velskBarony struct {
	feud.BaseBarony
}

var BKuloy_velsk库洛伊 feud.Barony = &库洛伊Kuloy_velskBarony{}

func init() {
    f := BKuloy_velsk库洛伊.(*库洛伊Kuloy_velskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuloy_velsk",
		TitleName: "库洛伊",
		TitleCode: "b_kuloy_velsk",
	}
}
