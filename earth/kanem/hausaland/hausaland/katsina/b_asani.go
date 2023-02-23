package katsina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿沙尼AsaniBarony struct {
	feud.BaseBarony
}

var BAsani阿沙尼 feud.Barony = &阿沙尼AsaniBarony{}

func init() {
	f := BAsani阿沙尼.(*阿沙尼AsaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asani",
		TitleName: "阿沙尼",
		TitleCode: "b_asani",
	}
}
