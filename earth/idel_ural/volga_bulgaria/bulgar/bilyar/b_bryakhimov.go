package bilyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布里亚希莫夫BryakhimovBarony struct {
	feud.BaseBarony
}

var BBryakhimov布里亚希莫夫 feud.Barony = &布里亚希莫夫BryakhimovBarony{}

func init() {
    f := BBryakhimov布里亚希莫夫.(*布里亚希莫夫BryakhimovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bryakhimov",
		TitleName: "布里亚希莫夫",
		TitleCode: "b_bryakhimov",
	}
}
