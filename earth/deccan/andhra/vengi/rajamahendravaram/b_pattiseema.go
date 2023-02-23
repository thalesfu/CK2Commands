package rajamahendravaram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕迪瑟玛PattiseemaBarony struct {
	feud.BaseBarony
}

var BPattiseema帕迪瑟玛 feud.Barony = &帕迪瑟玛PattiseemaBarony{}

func init() {
	f := BPattiseema帕迪瑟玛.(*帕迪瑟玛PattiseemaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pattiseema",
		TitleName: "帕迪瑟玛",
		TitleCode: "b_pattiseema",
	}
}
