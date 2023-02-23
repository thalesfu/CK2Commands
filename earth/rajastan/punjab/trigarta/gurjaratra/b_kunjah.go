package gurjaratra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡迦KunjahBarony struct {
	feud.BaseBarony
}

var BKunjah贡迦 feud.Barony = &贡迦KunjahBarony{}

func init() {
	f := BKunjah贡迦.(*贡迦KunjahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kunjah",
		TitleName: "贡迦",
		TitleCode: "b_kunjah",
	}
}
