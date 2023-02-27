package stettin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯德丁StettinBarony struct {
	feud.BaseBarony
}

var BStettin斯德丁 feud.Barony = &斯德丁StettinBarony{}

func init() {
    f := BStettin斯德丁.(*斯德丁StettinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stettin",
		TitleName: "斯德丁",
		TitleCode: "b_stettin",
	}
}
