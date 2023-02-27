package cebta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 休达CuetaBarony struct {
	feud.BaseBarony
}

var BCueta休达 feud.Barony = &休达CuetaBarony{}

func init() {
    f := BCueta休达.(*休达CuetaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cueta",
		TitleName: "休达",
		TitleCode: "b_cueta",
	}
}
