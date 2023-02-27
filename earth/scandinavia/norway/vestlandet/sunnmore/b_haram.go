package sunnmore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈拉姆HaramBarony struct {
	feud.BaseBarony
}

var BHaram哈拉姆 feud.Barony = &哈拉姆HaramBarony{}

func init() {
    f := BHaram哈拉姆.(*哈拉姆HaramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haram",
		TitleName: "哈拉姆",
		TitleCode: "b_haram",
	}
}
