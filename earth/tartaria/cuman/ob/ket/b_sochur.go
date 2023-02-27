package ket

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索丘尔SochurBarony struct {
	feud.BaseBarony
}

var BSochur索丘尔 feud.Barony = &索丘尔SochurBarony{}

func init() {
    f := BSochur索丘尔.(*索丘尔SochurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sochur",
		TitleName: "索丘尔",
		TitleCode: "b_sochur",
	}
}
