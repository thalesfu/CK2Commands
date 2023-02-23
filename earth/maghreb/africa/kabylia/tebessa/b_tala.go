package tebessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔拉TalaBarony struct {
	feud.BaseBarony
}

var BTala塔拉 feud.Barony = &塔拉TalaBarony{}

func init() {
	f := BTala塔拉.(*塔拉TalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tala",
		TitleName: "塔拉",
		TitleCode: "b_tala",
	}
}
