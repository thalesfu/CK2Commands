package figuig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加纳斯泰尔GanastelBarony struct {
	feud.BaseBarony
}

var BGanastel加纳斯泰尔 feud.Barony = &加纳斯泰尔GanastelBarony{}

func init() {
	f := BGanastel加纳斯泰尔.(*加纳斯泰尔GanastelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ganastel",
		TitleName: "加纳斯泰尔",
		TitleCode: "b_ganastel",
	}
}
