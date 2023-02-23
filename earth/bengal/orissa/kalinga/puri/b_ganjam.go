package puri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘贾姆GanjamBarony struct {
	feud.BaseBarony
}

var BGanjam甘贾姆 feud.Barony = &甘贾姆GanjamBarony{}

func init() {
	f := BGanjam甘贾姆.(*甘贾姆GanjamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ganjam",
		TitleName: "甘贾姆",
		TitleCode: "b_ganjam",
	}
}
