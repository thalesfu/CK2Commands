package prayaga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 钵罗耶伽PrayagaBarony struct {
	feud.BaseBarony
}

var BPrayaga钵罗耶伽 feud.Barony = &钵罗耶伽PrayagaBarony{}

func init() {
    f := BPrayaga钵罗耶伽.(*钵罗耶伽PrayagaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prayaga",
		TitleName: "钵罗耶伽",
		TitleCode: "b_prayaga",
	}
}
