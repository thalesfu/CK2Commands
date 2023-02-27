package taza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖尔西夫GeurcifBarony struct {
	feud.BaseBarony
}

var BGeurcif盖尔西夫 feud.Barony = &盖尔西夫GeurcifBarony{}

func init() {
    f := BGeurcif盖尔西夫.(*盖尔西夫GeurcifBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "geurcif",
		TitleName: "盖尔西夫",
		TitleCode: "b_geurcif",
	}
}
