package kharibta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯刻提斯ScetisBarony struct {
	feud.BaseBarony
}

var BScetis斯刻提斯 feud.Barony = &斯刻提斯ScetisBarony{}

func init() {
	f := BScetis斯刻提斯.(*斯刻提斯ScetisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "scetis",
		TitleName: "斯刻提斯",
		TitleCode: "b_scetis",
	}
}
