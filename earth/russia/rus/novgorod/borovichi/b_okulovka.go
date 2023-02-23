package borovichi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥库洛夫卡OkulovkaBarony struct {
	feud.BaseBarony
}

var BOkulovka奥库洛夫卡 feud.Barony = &奥库洛夫卡OkulovkaBarony{}

func init() {
	f := BOkulovka奥库洛夫卡.(*奥库洛夫卡OkulovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "okulovka",
		TitleName: "奥库洛夫卡",
		TitleCode: "b_okulovka",
	}
}
