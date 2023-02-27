package sripatha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿班和伽尼AbanhagariBarony struct {
	feud.BaseBarony
}

var BAbanhagari阿班和伽尼 feud.Barony = &阿班和伽尼AbanhagariBarony{}

func init() {
    f := BAbanhagari阿班和伽尼.(*阿班和伽尼AbanhagariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abanhagari",
		TitleName: "阿班和伽尼",
		TitleCode: "b_abanhagari",
	}
}
