package trondelag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯特洛特AustrattBarony struct {
	feud.BaseBarony
}

var BAustratt奥斯特洛特 feud.Barony = &奥斯特洛特AustrattBarony{}

func init() {
	f := BAustratt奥斯特洛特.(*奥斯特洛特AustrattBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "austratt",
		TitleName: "奥斯特洛特",
		TitleCode: "b_austratt",
	}
}
