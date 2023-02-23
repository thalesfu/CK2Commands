package uch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德拉瓦DerawarBarony struct {
	feud.BaseBarony
}

var BDerawar德拉瓦 feud.Barony = &德拉瓦DerawarBarony{}

func init() {
	f := BDerawar德拉瓦.(*德拉瓦DerawarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "derawar",
		TitleName: "德拉瓦",
		TitleCode: "b_derawar",
	}
}
