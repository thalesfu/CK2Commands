package almaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 新源XinyuanBarony struct {
	feud.BaseBarony
}

var BXinyuan新源 feud.Barony = &新源XinyuanBarony{}

func init() {
    f := BXinyuan新源.(*新源XinyuanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xinyuan",
		TitleName: "新源",
		TitleCode: "b_xinyuan",
	}
}
