package stezycka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 武库夫地区斯托切克StoczeklukowskiBarony struct {
	feud.BaseBarony
}

var BStoczeklukowski武库夫地区斯托切克 feud.Barony = &武库夫地区斯托切克StoczeklukowskiBarony{}

func init() {
    f := BStoczeklukowski武库夫地区斯托切克.(*武库夫地区斯托切克StoczeklukowskiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stoczeklukowski",
		TitleName: "武库夫地区斯托切克",
		TitleCode: "b_stoczeklukowski",
	}
}
