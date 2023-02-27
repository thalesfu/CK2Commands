package iona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托伯莫里TobermoryBarony struct {
	feud.BaseBarony
}

var BTobermory托伯莫里 feud.Barony = &托伯莫里TobermoryBarony{}

func init() {
    f := BTobermory托伯莫里.(*托伯莫里TobermoryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tobermory",
		TitleName: "托伯莫里",
		TitleCode: "b_tobermory",
	}
}
