package vaud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格吕耶尔GruyereBarony struct {
	feud.BaseBarony
}

var BGruyere格吕耶尔 feud.Barony = &格吕耶尔GruyereBarony{}

func init() {
    f := BGruyere格吕耶尔.(*格吕耶尔GruyereBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gruyere",
		TitleName: "格吕耶尔",
		TitleCode: "b_gruyere",
	}
}
