package asyut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 万尼亚WanninaBarony struct {
	feud.BaseBarony
}

var BWannina万尼亚 feud.Barony = &万尼亚WanninaBarony{}

func init() {
    f := BWannina万尼亚.(*万尼亚WanninaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wannina",
		TitleName: "万尼亚",
		TitleCode: "b_wannina",
	}
}
