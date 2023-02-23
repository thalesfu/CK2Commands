package tummana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩你因陀罗姞利呬ManendragarhBarony struct {
	feud.BaseBarony
}

var BManendragarh摩你因陀罗姞利呬 feud.Barony = &摩你因陀罗姞利呬ManendragarhBarony{}

func init() {
	f := BManendragarh摩你因陀罗姞利呬.(*摩你因陀罗姞利呬ManendragarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manendragarh",
		TitleName: "摩你因陀罗姞利呬",
		TitleCode: "b_manendragarh",
	}
}
