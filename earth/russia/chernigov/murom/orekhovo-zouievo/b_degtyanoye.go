package orekhovo-zouievo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰格佳诺耶DegtyanoyeBarony struct {
	feud.BaseBarony
}

var BDegtyanoye杰格佳诺耶 feud.Barony = &杰格佳诺耶DegtyanoyeBarony{}

func init() {
    f := BDegtyanoye杰格佳诺耶.(*杰格佳诺耶DegtyanoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "degtyanoye",
		TitleName: "杰格佳诺耶",
		TitleCode: "b_degtyanoye",
	}
}
