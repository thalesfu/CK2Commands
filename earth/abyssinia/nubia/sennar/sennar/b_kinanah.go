package sennar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉娜赫KinanahBarony struct {
	feud.BaseBarony
}

var BKinanah吉娜赫 feud.Barony = &吉娜赫KinanahBarony{}

func init() {
	f := BKinanah吉娜赫.(*吉娜赫KinanahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kinanah",
		TitleName: "吉娜赫",
		TitleCode: "b_kinanah",
	}
}
