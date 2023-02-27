package medelpad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥特曼OtmanBarony struct {
	feud.BaseBarony
}

var BOtman奥特曼 feud.Barony = &奥特曼OtmanBarony{}

func init() {
    f := BOtman奥特曼.(*奥特曼OtmanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "otman",
		TitleName: "奥特曼",
		TitleCode: "b_otman",
	}
}
