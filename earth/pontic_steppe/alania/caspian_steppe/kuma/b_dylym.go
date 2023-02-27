package kuma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德雷姆DylymBarony struct {
	feud.BaseBarony
}

var BDylym德雷姆 feud.Barony = &德雷姆DylymBarony{}

func init() {
    f := BDylym德雷姆.(*德雷姆DylymBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dylym",
		TitleName: "德雷姆",
		TitleCode: "b_dylym",
	}
}
