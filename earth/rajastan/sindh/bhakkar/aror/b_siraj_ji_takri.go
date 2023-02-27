package aror

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯利加_吉_塔克里Siraj_ji_takriBarony struct {
	feud.BaseBarony
}

var BSiraj_ji_takri斯利加_吉_塔克里 feud.Barony = &斯利加_吉_塔克里Siraj_ji_takriBarony{}

func init() {
    f := BSiraj_ji_takri斯利加_吉_塔克里.(*斯利加_吉_塔克里Siraj_ji_takriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siraj_ji_takri",
		TitleName: "斯利加_吉_塔克里",
		TitleCode: "b_siraj_ji_takri",
	}
}
