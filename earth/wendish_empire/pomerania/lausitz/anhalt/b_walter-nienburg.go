package anhalt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔特宁堡Walter_nienburgBarony struct {
	feud.BaseBarony
}

var BWalter_nienburg瓦尔特宁堡 feud.Barony = &瓦尔特宁堡Walter_nienburgBarony{}

func init() {
    f := BWalter_nienburg瓦尔特宁堡.(*瓦尔特宁堡Walter_nienburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "walter_nienburg",
		TitleName: "瓦尔特宁堡",
		TitleCode: "b_walter-nienburg",
	}
}
