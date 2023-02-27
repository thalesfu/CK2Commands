package asyut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜鲁卡DurunkaBarony struct {
	feud.BaseBarony
}

var BDurunka杜鲁卡 feud.Barony = &杜鲁卡DurunkaBarony{}

func init() {
    f := BDurunka杜鲁卡.(*杜鲁卡DurunkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durunka",
		TitleName: "杜鲁卡",
		TitleCode: "b_durunka",
	}
}
