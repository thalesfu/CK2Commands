package bome

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那哈NahaBarony struct {
	feud.BaseBarony
}

var BNaha那哈 feud.Barony = &那哈NahaBarony{}

func init() {
    f := BNaha那哈.(*那哈NahaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naha",
		TitleName: "那哈",
		TitleCode: "b_naha",
	}
}
