package mosul

import (
	"github.com/thalesfu/CK2Commands/earth/persia/mesopotamia/mosul/mosul"
	"github.com/thalesfu/CK2Commands/earth/persia/mesopotamia/mosul/nisibin"
	"github.com/thalesfu/CK2Commands/earth/persia/mesopotamia/mosul/oromieh"
	"github.com/thalesfu/CK2Commands/earth/persia/mesopotamia/mosul/sinjar"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MosulDuke interface {
    feud.Duke
    CMosul摩苏尔() 	mosul.MosulCounty
    CNisibin尼西宾() 	nisibin.NisibinCounty
    COromieh埃尔比勒() 	oromieh.OromiehCounty
    CSinjar辛贾尔() 	sinjar.SinjarCounty
}

type 摩苏尔MosulDuke struct {
	feud.BaseDuke
	摩苏尔Mosul 	mosul.MosulCounty
	尼西宾Nisibin 	nisibin.NisibinCounty
	埃尔比勒Oromieh 	oromieh.OromiehCounty
	辛贾尔Sinjar 	sinjar.SinjarCounty
}

func (d *摩苏尔MosulDuke) CMosul摩苏尔() mosul.MosulCounty {
	return d.摩苏尔Mosul
}
    
func (d *摩苏尔MosulDuke) CNisibin尼西宾() nisibin.NisibinCounty {
	return d.尼西宾Nisibin
}
    
func (d *摩苏尔MosulDuke) COromieh埃尔比勒() oromieh.OromiehCounty {
	return d.埃尔比勒Oromieh
}
    
func (d *摩苏尔MosulDuke) CSinjar辛贾尔() sinjar.SinjarCounty {
	return d.辛贾尔Sinjar
}
    
var DMosul摩苏尔 MosulDuke = &摩苏尔MosulDuke{}

func init() {
	f := DMosul摩苏尔.(*摩苏尔MosulDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "mosul",
		TitleName: "摩苏尔",
		TitleCode: "d_mosul",
		Counties:  map[string]feud.County{},
	}

	f.摩苏尔Mosul = mosul.CMosul摩苏尔
	f.摩苏尔Mosul.SetParent(f)
	
	f.尼西宾Nisibin = nisibin.CNisibin尼西宾
	f.尼西宾Nisibin.SetParent(f)
	
	f.埃尔比勒Oromieh = oromieh.COromieh埃尔比勒
	f.埃尔比勒Oromieh.SetParent(f)
	
	f.辛贾尔Sinjar = sinjar.CSinjar辛贾尔
	f.辛贾尔Sinjar.SetParent(f)
	
}
