package kuopio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库卡林湖KukkarinselkaBarony struct {
	feud.BaseBarony
}

var BKukkarinselka库卡林湖 feud.Barony = &库卡林湖KukkarinselkaBarony{}

func init() {
    f := BKukkarinselka库卡林湖.(*库卡林湖KukkarinselkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kukkarinselka",
		TitleName: "库卡林湖",
		TitleCode: "b_kukkarinselka",
	}
}
