package ankyra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰尔玛GermaBarony struct {
	feud.BaseBarony
}

var BGerma杰尔玛 feud.Barony = &杰尔玛GermaBarony{}

func init() {
    f := BGerma杰尔玛.(*杰尔玛GermaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "germa",
		TitleName: "杰尔玛",
		TitleCode: "b_germa",
	}
}
