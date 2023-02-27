package bilyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌鲁苏UrussuBarony struct {
	feud.BaseBarony
}

var BUrussu乌鲁苏 feud.Barony = &乌鲁苏UrussuBarony{}

func init() {
    f := BUrussu乌鲁苏.(*乌鲁苏UrussuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urussu",
		TitleName: "乌鲁苏",
		TitleCode: "b_urussu",
	}
}
