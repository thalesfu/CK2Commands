package beloozero

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费拉蓬托夫FerapontovBarony struct {
	feud.BaseBarony
}

var BFerapontov费拉蓬托夫 feud.Barony = &费拉蓬托夫FerapontovBarony{}

func init() {
    f := BFerapontov费拉蓬托夫.(*费拉蓬托夫FerapontovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ferapontov",
		TitleName: "费拉蓬托夫",
		TitleCode: "b_ferapontov",
	}
}
