package chauragarh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯娑梨SirsaliBarony struct {
	feud.BaseBarony
}

var BSirsali斯娑梨 feud.Barony = &斯娑梨SirsaliBarony{}

func init() {
	f := BSirsali斯娑梨.(*斯娑梨SirsaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sirsali",
		TitleName: "斯娑梨",
		TitleCode: "b_sirsali",
	}
}
