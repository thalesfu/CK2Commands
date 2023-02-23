package buqtirma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布赫塔尔马BuqtirmaBarony struct {
	feud.BaseBarony
}

var BBuqtirma布赫塔尔马 feud.Barony = &布赫塔尔马BuqtirmaBarony{}

func init() {
	f := BBuqtirma布赫塔尔马.(*布赫塔尔马BuqtirmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buqtirma",
		TitleName: "布赫塔尔马",
		TitleCode: "b_buqtirma",
	}
}
