package kanyakubja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇耶旃陀罗Jai_chandraBarony struct {
	feud.BaseBarony
}

var BJai_chandra阇耶旃陀罗 feud.Barony = &阇耶旃陀罗Jai_chandraBarony{}

func init() {
    f := BJai_chandra阇耶旃陀罗.(*阇耶旃陀罗Jai_chandraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jai_chandra",
		TitleName: "阇耶旃陀罗",
		TitleCode: "b_jai_chandra",
	}
}
