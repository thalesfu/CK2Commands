package syktyvkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟索拉SysolaBarony struct {
	feud.BaseBarony
}

var BSysola瑟索拉 feud.Barony = &瑟索拉SysolaBarony{}

func init() {
    f := BSysola瑟索拉.(*瑟索拉SysolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sysola",
		TitleName: "瑟索拉",
		TitleCode: "b_sysola",
	}
}
