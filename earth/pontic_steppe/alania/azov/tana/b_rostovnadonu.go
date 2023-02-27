package tana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 顿河畔罗斯托夫RostovnadonuBarony struct {
	feud.BaseBarony
}

var BRostovnadonu顿河畔罗斯托夫 feud.Barony = &顿河畔罗斯托夫RostovnadonuBarony{}

func init() {
    f := BRostovnadonu顿河畔罗斯托夫.(*顿河畔罗斯托夫RostovnadonuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rostovnadonu",
		TitleName: "顿河畔罗斯托夫",
		TitleCode: "b_rostovnadonu",
	}
}
