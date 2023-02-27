package chios

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕贡达斯PagondasBarony struct {
	feud.BaseBarony
}

var BPagondas帕贡达斯 feud.Barony = &帕贡达斯PagondasBarony{}

func init() {
    f := BPagondas帕贡达斯.(*帕贡达斯PagondasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pagondas",
		TitleName: "帕贡达斯",
		TitleCode: "b_pagondas",
	}
}
