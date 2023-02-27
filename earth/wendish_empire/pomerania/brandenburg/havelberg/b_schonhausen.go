package havelberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 申豪森SchonhausenBarony struct {
	feud.BaseBarony
}

var BSchonhausen申豪森 feud.Barony = &申豪森SchonhausenBarony{}

func init() {
    f := BSchonhausen申豪森.(*申豪森SchonhausenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schonhausen",
		TitleName: "申豪森",
		TitleCode: "b_schonhausen",
	}
}
