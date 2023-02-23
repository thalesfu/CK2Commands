package sancerre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑塞尔SancerreBarony struct {
	feud.BaseBarony
}

var BSancerre桑塞尔 feud.Barony = &桑塞尔SancerreBarony{}

func init() {
	f := BSancerre桑塞尔.(*桑塞尔SancerreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sancerre",
		TitleName: "桑塞尔",
		TitleCode: "b_sancerre",
	}
}
