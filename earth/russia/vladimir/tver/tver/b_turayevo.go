package tver

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图拉耶沃TurayevoBarony struct {
	feud.BaseBarony
}

var BTurayevo图拉耶沃 feud.Barony = &图拉耶沃TurayevoBarony{}

func init() {
	f := BTurayevo图拉耶沃.(*图拉耶沃TurayevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turayevo",
		TitleName: "图拉耶沃",
		TitleCode: "b_turayevo",
	}
}
