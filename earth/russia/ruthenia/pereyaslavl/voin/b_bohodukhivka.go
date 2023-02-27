package voin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博戈杜霍夫卡BohodukhivkaBarony struct {
	feud.BaseBarony
}

var BBohodukhivka博戈杜霍夫卡 feud.Barony = &博戈杜霍夫卡BohodukhivkaBarony{}

func init() {
    f := BBohodukhivka博戈杜霍夫卡.(*博戈杜霍夫卡BohodukhivkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bohodukhivka",
		TitleName: "博戈杜霍夫卡",
		TitleCode: "b_bohodukhivka",
	}
}
