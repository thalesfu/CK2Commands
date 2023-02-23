package merv

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 木鹿MervBarony struct {
	feud.BaseBarony
}

var BMerv木鹿 feud.Barony = &木鹿MervBarony{}

func init() {
	f := BMerv木鹿.(*木鹿MervBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "merv",
		TitleName: "木鹿",
		TitleCode: "b_merv",
	}
}
