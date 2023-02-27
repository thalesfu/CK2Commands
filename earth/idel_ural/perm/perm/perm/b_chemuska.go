package perm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切穆斯卡ChemuskaBarony struct {
	feud.BaseBarony
}

var BChemuska切穆斯卡 feud.Barony = &切穆斯卡ChemuskaBarony{}

func init() {
    f := BChemuska切穆斯卡.(*切穆斯卡ChemuskaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chemuska",
		TitleName: "切穆斯卡",
		TitleCode: "b_chemuska",
	}
}
