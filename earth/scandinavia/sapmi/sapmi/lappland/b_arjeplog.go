package lappland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔耶普卢格ArjeplogBarony struct {
	feud.BaseBarony
}

var BArjeplog阿尔耶普卢格 feud.Barony = &阿尔耶普卢格ArjeplogBarony{}

func init() {
    f := BArjeplog阿尔耶普卢格.(*阿尔耶普卢格ArjeplogBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arjeplog",
		TitleName: "阿尔耶普卢格",
		TitleCode: "b_arjeplog",
	}
}
