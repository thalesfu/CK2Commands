package leinster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克洛ArklowBarony struct {
	feud.BaseBarony
}

var BArklow阿克洛 feud.Barony = &阿克洛ArklowBarony{}

func init() {
    f := BArklow阿克洛.(*阿克洛ArklowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arklow",
		TitleName: "阿克洛",
		TitleCode: "b_arklow",
	}
}
