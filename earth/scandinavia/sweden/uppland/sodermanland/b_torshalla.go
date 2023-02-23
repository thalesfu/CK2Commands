package sodermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托什海拉TorshallaBarony struct {
	feud.BaseBarony
}

var BTorshalla托什海拉 feud.Barony = &托什海拉TorshallaBarony{}

func init() {
	f := BTorshalla托什海拉.(*托什海拉TorshallaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torshalla",
		TitleName: "托什海拉",
		TitleCode: "b_torshalla",
	}
}
