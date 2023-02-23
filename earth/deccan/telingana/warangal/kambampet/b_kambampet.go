package kambampet

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 剑婆摩波底KambampetBarony struct {
	feud.BaseBarony
}

var BKambampet剑婆摩波底 feud.Barony = &剑婆摩波底KambampetBarony{}

func init() {
	f := BKambampet剑婆摩波底.(*剑婆摩波底KambampetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kambampet",
		TitleName: "剑婆摩波底",
		TitleCode: "b_kambampet",
	}
}
