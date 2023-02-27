package lukomorie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古利艾波列HuliaipoleBarony struct {
	feud.BaseBarony
}

var BHuliaipole古利艾波列 feud.Barony = &古利艾波列HuliaipoleBarony{}

func init() {
    f := BHuliaipole古利艾波列.(*古利艾波列HuliaipoleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huliaipole",
		TitleName: "古利艾波列",
		TitleCode: "b_huliaipole",
	}
}
