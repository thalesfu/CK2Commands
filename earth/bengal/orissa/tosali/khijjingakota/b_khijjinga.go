package khijjingakota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 契承伽KhijjingaBarony struct {
	feud.BaseBarony
}

var BKhijjinga契承伽 feud.Barony = &契承伽KhijjingaBarony{}

func init() {
    f := BKhijjinga契承伽.(*契承伽KhijjingaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khijjinga",
		TitleName: "契承伽",
		TitleCode: "b_khijjinga",
	}
}
