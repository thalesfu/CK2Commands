package rottenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 魏恩斯贝格WeinsbergBarony struct {
	feud.BaseBarony
}

var BWeinsberg魏恩斯贝格 feud.Barony = &魏恩斯贝格WeinsbergBarony{}

func init() {
    f := BWeinsberg魏恩斯贝格.(*魏恩斯贝格WeinsbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "weinsberg",
		TitleName: "魏恩斯贝格",
		TitleCode: "b_weinsberg",
	}
}
