package halland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔贝里VarbergBarony struct {
	feud.BaseBarony
}

var BVarberg瓦尔贝里 feud.Barony = &瓦尔贝里VarbergBarony{}

func init() {
    f := BVarberg瓦尔贝里.(*瓦尔贝里VarbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varberg",
		TitleName: "瓦尔贝里",
		TitleCode: "b_varberg",
	}
}
