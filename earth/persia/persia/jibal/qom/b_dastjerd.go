package qom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达斯特杰尔德DastjerdBarony struct {
	feud.BaseBarony
}

var BDastjerd达斯特杰尔德 feud.Barony = &达斯特杰尔德DastjerdBarony{}

func init() {
    f := BDastjerd达斯特杰尔德.(*达斯特杰尔德DastjerdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dastjerd",
		TitleName: "达斯特杰尔德",
		TitleCode: "b_dastjerd",
	}
}
