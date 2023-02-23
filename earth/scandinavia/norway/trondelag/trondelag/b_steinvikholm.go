package trondelag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯泰因维克霍姆SteinvikholmBarony struct {
	feud.BaseBarony
}

var BSteinvikholm斯泰因维克霍姆 feud.Barony = &斯泰因维克霍姆SteinvikholmBarony{}

func init() {
	f := BSteinvikholm斯泰因维克霍姆.(*斯泰因维克霍姆SteinvikholmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "steinvikholm",
		TitleName: "斯泰因维克霍姆",
		TitleCode: "b_steinvikholm",
	}
}
