package kolguyev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布格良卡BugryankaBarony struct {
	feud.BaseBarony
}

var BBugryanka布格良卡 feud.Barony = &布格良卡BugryankaBarony{}

func init() {
    f := BBugryanka布格良卡.(*布格良卡BugryankaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bugryanka",
		TitleName: "布格良卡",
		TitleCode: "b_bugryanka",
	}
}
