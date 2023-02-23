package sunnmore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯泰因沃格SteinvagBarony struct {
	feud.BaseBarony
}

var BSteinvag斯泰因沃格 feud.Barony = &斯泰因沃格SteinvagBarony{}

func init() {
	f := BSteinvag斯泰因沃格.(*斯泰因沃格SteinvagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "steinvag",
		TitleName: "斯泰因沃格",
		TitleCode: "b_steinvag",
	}
}
