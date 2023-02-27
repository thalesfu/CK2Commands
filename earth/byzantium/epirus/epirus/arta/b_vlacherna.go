package arta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗拉海尔纳VlachernaBarony struct {
	feud.BaseBarony
}

var BVlacherna弗拉海尔纳 feud.Barony = &弗拉海尔纳VlachernaBarony{}

func init() {
    f := BVlacherna弗拉海尔纳.(*弗拉海尔纳VlachernaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vlacherna",
		TitleName: "弗拉海尔纳",
		TitleCode: "b_vlacherna",
	}
}
