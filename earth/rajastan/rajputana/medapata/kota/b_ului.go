package kota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌卢伊UluiBarony struct {
	feud.BaseBarony
}

var BUlui乌卢伊 feud.Barony = &乌卢伊UluiBarony{}

func init() {
	f := BUlui乌卢伊.(*乌卢伊UluiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ului",
		TitleName: "乌卢伊",
		TitleCode: "b_ului",
	}
}
