package tell_bashir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别尔哈BirthaBarony struct {
	feud.BaseBarony
}

var BBirtha别尔哈 feud.Barony = &别尔哈BirthaBarony{}

func init() {
    f := BBirtha别尔哈.(*别尔哈BirthaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "birtha",
		TitleName: "别尔哈",
		TitleCode: "b_birtha",
	}
}
