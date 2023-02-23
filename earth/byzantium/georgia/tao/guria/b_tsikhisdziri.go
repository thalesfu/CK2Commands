package guria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 齐希斯济里TsikhisdziriBarony struct {
	feud.BaseBarony
}

var BTsikhisdziri齐希斯济里 feud.Barony = &齐希斯济里TsikhisdziriBarony{}

func init() {
	f := BTsikhisdziri齐希斯济里.(*齐希斯济里TsikhisdziriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsikhisdziri",
		TitleName: "齐希斯济里",
		TitleCode: "b_tsikhisdziri",
	}
}
