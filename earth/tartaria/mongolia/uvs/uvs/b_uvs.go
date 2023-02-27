package uvs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌布苏UvsBarony struct {
	feud.BaseBarony
}

var BUvs乌布苏 feud.Barony = &乌布苏UvsBarony{}

func init() {
    f := BUvs乌布苏.(*乌布苏UvsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uvs",
		TitleName: "乌布苏",
		TitleCode: "b_uvs",
	}
}
