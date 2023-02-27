package plzen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普拉西PlasyBarony struct {
	feud.BaseBarony
}

var BPlasy普拉西 feud.Barony = &普拉西PlasyBarony{}

func init() {
    f := BPlasy普拉西.(*普拉西PlasyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plasy",
		TitleName: "普拉西",
		TitleCode: "b_plasy",
	}
}
