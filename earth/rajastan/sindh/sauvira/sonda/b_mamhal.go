package sonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙诃尔MamhalBarony struct {
	feud.BaseBarony
}

var BMamhal蒙诃尔 feud.Barony = &蒙诃尔MamhalBarony{}

func init() {
	f := BMamhal蒙诃尔.(*蒙诃尔MamhalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mamhal",
		TitleName: "蒙诃尔",
		TitleCode: "b_mamhal",
	}
}
