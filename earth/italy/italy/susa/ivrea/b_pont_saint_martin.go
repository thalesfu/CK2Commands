package ivrea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蓬圣马丁Pont_saint_martinBarony struct {
	feud.BaseBarony
}

var BPont_saint_martin蓬圣马丁 feud.Barony = &蓬圣马丁Pont_saint_martinBarony{}

func init() {
    f := BPont_saint_martin蓬圣马丁.(*蓬圣马丁Pont_saint_martinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pont_saint_martin",
		TitleName: "蓬圣马丁",
		TitleCode: "b_pont_saint_martin",
	}
}
