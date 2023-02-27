package balanjar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔贡ArgunBarony struct {
	feud.BaseBarony
}

var BArgun阿尔贡 feud.Barony = &阿尔贡ArgunBarony{}

func init() {
    f := BArgun阿尔贡.(*阿尔贡ArgunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "argun",
		TitleName: "阿尔贡",
		TitleCode: "b_argun",
	}
}
