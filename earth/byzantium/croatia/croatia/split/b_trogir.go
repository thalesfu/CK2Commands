package split

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特罗吉尔TrogirBarony struct {
	feud.BaseBarony
}

var BTrogir特罗吉尔 feud.Barony = &特罗吉尔TrogirBarony{}

func init() {
	f := BTrogir特罗吉尔.(*特罗吉尔TrogirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trogir",
		TitleName: "特罗吉尔",
		TitleCode: "b_trogir",
	}
}
