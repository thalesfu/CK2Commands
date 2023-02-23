package suvarnagram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩醯因陀罗姞利呬MahendragarhBarony struct {
	feud.BaseBarony
}

var BMahendragarh摩醯因陀罗姞利呬 feud.Barony = &摩醯因陀罗姞利呬MahendragarhBarony{}

func init() {
	f := BMahendragarh摩醯因陀罗姞利呬.(*摩醯因陀罗姞利呬MahendragarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahendragarh",
		TitleName: "摩醯因陀罗姞利呬",
		TitleCode: "b_mahendragarh",
	}
}
