package derbent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰尔宾特DerbentBarony struct {
	feud.BaseBarony
}

var BDerbent杰尔宾特 feud.Barony = &杰尔宾特DerbentBarony{}

func init() {
    f := BDerbent杰尔宾特.(*杰尔宾特DerbentBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "derbent",
		TitleName: "杰尔宾特",
		TitleCode: "b_derbent",
	}
}
