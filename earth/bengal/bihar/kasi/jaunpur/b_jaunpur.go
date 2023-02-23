package jaunpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 善补罗JaunpurBarony struct {
	feud.BaseBarony
}

var BJaunpur善补罗 feud.Barony = &善补罗JaunpurBarony{}

func init() {
	f := BJaunpur善补罗.(*善补罗JaunpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jaunpur",
		TitleName: "善补罗",
		TitleCode: "b_jaunpur",
	}
}
