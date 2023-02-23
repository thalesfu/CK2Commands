package aleppo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古赛尔QusayrBarony struct {
	feud.BaseBarony
}

var BQusayr古赛尔 feud.Barony = &古赛尔QusayrBarony{}

func init() {
	f := BQusayr古赛尔.(*古赛尔QusayrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qusayr",
		TitleName: "古赛尔",
		TitleCode: "b_qusayr",
	}
}
