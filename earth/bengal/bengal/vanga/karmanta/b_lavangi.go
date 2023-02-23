package karmanta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢槃吉LavangiBarony struct {
	feud.BaseBarony
}

var BLavangi卢槃吉 feud.Barony = &卢槃吉LavangiBarony{}

func init() {
	f := BLavangi卢槃吉.(*卢槃吉LavangiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lavangi",
		TitleName: "卢槃吉",
		TitleCode: "b_lavangi",
	}
}
