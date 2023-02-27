package kakheti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰拉维TelaviBarony struct {
	feud.BaseBarony
}

var BTelavi泰拉维 feud.Barony = &泰拉维TelaviBarony{}

func init() {
    f := BTelavi泰拉维.(*泰拉维TelaviBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "telavi",
		TitleName: "泰拉维",
		TitleCode: "b_telavi",
	}
}
