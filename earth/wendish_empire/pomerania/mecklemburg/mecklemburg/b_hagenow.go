package mecklemburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈格诺HagenowBarony struct {
	feud.BaseBarony
}

var BHagenow哈格诺 feud.Barony = &哈格诺HagenowBarony{}

func init() {
    f := BHagenow哈格诺.(*哈格诺HagenowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hagenow",
		TitleName: "哈格诺",
		TitleCode: "b_hagenow",
	}
}
