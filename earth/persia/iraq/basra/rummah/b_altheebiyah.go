package rummah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿特比杨AltheebiyahBarony struct {
	feud.BaseBarony
}

var BAltheebiyah阿特比杨 feud.Barony = &阿特比杨AltheebiyahBarony{}

func init() {
	f := BAltheebiyah阿特比杨.(*阿特比杨AltheebiyahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "altheebiyah",
		TitleName: "阿特比杨",
		TitleCode: "b_altheebiyah",
	}
}
