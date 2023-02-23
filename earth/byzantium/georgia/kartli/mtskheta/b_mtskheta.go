package mtskheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姆茨赫塔MtskhetaBarony struct {
	feud.BaseBarony
}

var BMtskheta姆茨赫塔 feud.Barony = &姆茨赫塔MtskhetaBarony{}

func init() {
	f := BMtskheta姆茨赫塔.(*姆茨赫塔MtskhetaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mtskheta",
		TitleName: "姆茨赫塔",
		TitleCode: "b_mtskheta",
	}
}
