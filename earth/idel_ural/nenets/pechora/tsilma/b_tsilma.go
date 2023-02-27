package tsilma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 齐利马TsilmaBarony struct {
	feud.BaseBarony
}

var BTsilma齐利马 feud.Barony = &齐利马TsilmaBarony{}

func init() {
    f := BTsilma齐利马.(*齐利马TsilmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsilma",
		TitleName: "齐利马",
		TitleCode: "b_tsilma",
	}
}
