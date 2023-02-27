package turov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯托林StolinBarony struct {
	feud.BaseBarony
}

var BStolin斯托林 feud.Barony = &斯托林StolinBarony{}

func init() {
    f := BStolin斯托林.(*斯托林StolinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stolin",
		TitleName: "斯托林",
		TitleCode: "b_stolin",
	}
}
