package orava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内迈特利普切NemetlipcseBarony struct {
	feud.BaseBarony
}

var BNemetlipcse内迈特利普切 feud.Barony = &内迈特利普切NemetlipcseBarony{}

func init() {
    f := BNemetlipcse内迈特利普切.(*内迈特利普切NemetlipcseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nemetlipcse",
		TitleName: "内迈特利普切",
		TitleCode: "b_nemetlipcse",
	}
}
