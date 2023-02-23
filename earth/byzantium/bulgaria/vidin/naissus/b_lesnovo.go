package naissus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱斯诺沃LesnovoBarony struct {
	feud.BaseBarony
}

var BLesnovo莱斯诺沃 feud.Barony = &莱斯诺沃LesnovoBarony{}

func init() {
	f := BLesnovo莱斯诺沃.(*莱斯诺沃LesnovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lesnovo",
		TitleName: "莱斯诺沃",
		TitleCode: "b_lesnovo",
	}
}
