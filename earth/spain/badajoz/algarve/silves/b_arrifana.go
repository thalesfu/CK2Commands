package silves

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿里法纳ArrifanaBarony struct {
	feud.BaseBarony
}

var BArrifana阿里法纳 feud.Barony = &阿里法纳ArrifanaBarony{}

func init() {
    f := BArrifana阿里法纳.(*阿里法纳ArrifanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arrifana",
		TitleName: "阿里法纳",
		TitleCode: "b_arrifana",
	}
}
