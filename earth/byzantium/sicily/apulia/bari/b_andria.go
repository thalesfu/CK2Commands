package bari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安德里亚AndriaBarony struct {
	feud.BaseBarony
}

var BAndria安德里亚 feud.Barony = &安德里亚AndriaBarony{}

func init() {
    f := BAndria安德里亚.(*安德里亚AndriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "andria",
		TitleName: "安德里亚",
		TitleCode: "b_andria",
	}
}
