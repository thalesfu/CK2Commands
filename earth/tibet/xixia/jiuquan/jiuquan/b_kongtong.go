package jiuquan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 崆峒KongtongBarony struct {
	feud.BaseBarony
}

var BKongtong崆峒 feud.Barony = &崆峒KongtongBarony{}

func init() {
    f := BKongtong崆峒.(*崆峒KongtongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kongtong",
		TitleName: "崆峒",
		TitleCode: "b_kongtong",
	}
}
