package stezycka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯滕日察StezycaBarony struct {
	feud.BaseBarony
}

var BStezyca斯滕日察 feud.Barony = &斯滕日察StezycaBarony{}

func init() {
    f := BStezyca斯滕日察.(*斯滕日察StezycaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stezyca",
		TitleName: "斯滕日察",
		TitleCode: "b_stezyca",
	}
}
