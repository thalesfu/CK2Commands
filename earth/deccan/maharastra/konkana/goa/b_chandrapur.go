package goa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旃陀罗补罗ChandrapurBarony struct {
	feud.BaseBarony
}

var BChandrapur旃陀罗补罗 feud.Barony = &旃陀罗补罗ChandrapurBarony{}

func init() {
	f := BChandrapur旃陀罗补罗.(*旃陀罗补罗ChandrapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chandrapur",
		TitleName: "旃陀罗补罗",
		TitleCode: "b_chandrapur",
	}
}
