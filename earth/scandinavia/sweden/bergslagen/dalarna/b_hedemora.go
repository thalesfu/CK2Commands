package dalarna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海德穆拉HedemoraBarony struct {
	feud.BaseBarony
}

var BHedemora海德穆拉 feud.Barony = &海德穆拉HedemoraBarony{}

func init() {
	f := BHedemora海德穆拉.(*海德穆拉HedemoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hedemora",
		TitleName: "海德穆拉",
		TitleCode: "b_hedemora",
	}
}
