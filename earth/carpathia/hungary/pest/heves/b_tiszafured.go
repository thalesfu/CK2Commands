package heves

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂绍菲赖德TiszafuredBarony struct {
	feud.BaseBarony
}

var BTiszafured蒂绍菲赖德 feud.Barony = &蒂绍菲赖德TiszafuredBarony{}

func init() {
	f := BTiszafured蒂绍菲赖德.(*蒂绍菲赖德TiszafuredBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiszafured",
		TitleName: "蒂绍菲赖德",
		TitleCode: "b_tiszafured",
	}
}
