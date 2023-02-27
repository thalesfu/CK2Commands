package agrigento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利卡塔LicataBarony struct {
	feud.BaseBarony
}

var BLicata利卡塔 feud.Barony = &利卡塔LicataBarony{}

func init() {
    f := BLicata利卡塔.(*利卡塔LicataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "licata",
		TitleName: "利卡塔",
		TitleCode: "b_licata",
	}
}
