package rajrappa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 豆富DhaurpurBarony struct {
	feud.BaseBarony
}

var BDhaurpur豆富 feud.Barony = &豆富DhaurpurBarony{}

func init() {
    f := BDhaurpur豆富.(*豆富DhaurpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhaurpur",
		TitleName: "豆富",
		TitleCode: "b_dhaurpur",
	}
}
