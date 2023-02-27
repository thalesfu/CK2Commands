package werle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施特拉尔松德StralsundBarony struct {
	feud.BaseBarony
}

var BStralsund施特拉尔松德 feud.Barony = &施特拉尔松德StralsundBarony{}

func init() {
    f := BStralsund施特拉尔松德.(*施特拉尔松德StralsundBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stralsund",
		TitleName: "施特拉尔松德",
		TitleCode: "b_stralsund",
	}
}
