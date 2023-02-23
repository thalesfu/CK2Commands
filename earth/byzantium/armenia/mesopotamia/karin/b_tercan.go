package karin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰尔詹TercanBarony struct {
	feud.BaseBarony
}

var BTercan泰尔詹 feud.Barony = &泰尔詹TercanBarony{}

func init() {
	f := BTercan泰尔詹.(*泰尔詹TercanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tercan",
		TitleName: "泰尔詹",
		TitleCode: "b_tercan",
	}
}
