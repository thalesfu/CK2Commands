package derby

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 威克斯沃斯WirksworthBarony struct {
	feud.BaseBarony
}

var BWirksworth威克斯沃斯 feud.Barony = &威克斯沃斯WirksworthBarony{}

func init() {
	f := BWirksworth威克斯沃斯.(*威克斯沃斯WirksworthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wirksworth",
		TitleName: "威克斯沃斯",
		TitleCode: "b_wirksworth",
	}
}
