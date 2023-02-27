package penugonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 具罗摩军荼GurramkondaBarony struct {
	feud.BaseBarony
}

var BGurramkonda具罗摩军荼 feud.Barony = &具罗摩军荼GurramkondaBarony{}

func init() {
    f := BGurramkonda具罗摩军荼.(*具罗摩军荼GurramkondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gurramkonda",
		TitleName: "具罗摩军荼",
		TitleCode: "b_gurramkonda",
	}
}
