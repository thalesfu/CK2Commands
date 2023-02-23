package medelpad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托普TorpBarony struct {
	feud.BaseBarony
}

var BTorp托普 feud.Barony = &托普TorpBarony{}

func init() {
	f := BTorp托普.(*托普TorpBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torp",
		TitleName: "托普",
		TitleCode: "b_torp",
	}
}
