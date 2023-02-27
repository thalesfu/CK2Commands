package koshma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克沙尔KesharBarony struct {
	feud.BaseBarony
}

var BKeshar克沙尔 feud.Barony = &克沙尔KesharBarony{}

func init() {
    f := BKeshar克沙尔.(*克沙尔KesharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "keshar",
		TitleName: "克沙尔",
		TitleCode: "b_keshar",
	}
}
