package kunlun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 独立石DulihriBarony struct {
	feud.BaseBarony
}

var BDulihri独立石 feud.Barony = &独立石DulihriBarony{}

func init() {
	f := BDulihri独立石.(*独立石DulihriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dulihri",
		TitleName: "独立石",
		TitleCode: "b_dulihri",
	}
}
