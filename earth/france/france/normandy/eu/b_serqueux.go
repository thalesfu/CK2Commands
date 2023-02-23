package eu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞尔克SerqueuxBarony struct {
	feud.BaseBarony
}

var BSerqueux塞尔克 feud.Barony = &塞尔克SerqueuxBarony{}

func init() {
	f := BSerqueux塞尔克.(*塞尔克SerqueuxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "serqueux",
		TitleName: "塞尔克",
		TitleCode: "b_serqueux",
	}
}
