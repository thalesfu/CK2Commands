package sagar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达利戈达TalikotaBarony struct {
	feud.BaseBarony
}

var BTalikota达利戈达 feud.Barony = &达利戈达TalikotaBarony{}

func init() {
	f := BTalikota达利戈达.(*达利戈达TalikotaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talikota",
		TitleName: "达利戈达",
		TitleCode: "b_talikota",
	}
}
