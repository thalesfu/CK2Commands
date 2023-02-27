package thalner

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 申陀契荼SindkhedaBarony struct {
	feud.BaseBarony
}

var BSindkheda申陀契荼 feud.Barony = &申陀契荼SindkhedaBarony{}

func init() {
    f := BSindkheda申陀契荼.(*申陀契荼SindkhedaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sindkheda",
		TitleName: "申陀契荼",
		TitleCode: "b_sindkheda",
	}
}
