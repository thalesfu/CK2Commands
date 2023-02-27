package rostov

import (
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/rostov/beloozero"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/rostov/rostov"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/rostov/uglich"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/rostov/yaroslavl"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RostovDuke interface {
    feud.Duke
    CBeloozero别洛奥泽罗() 	beloozero.BeloozeroCounty
    CRostov罗斯托夫() 	rostov.RostovCounty
    CUglich乌格利奇() 	uglich.UglichCounty
    CYaroslavl雅罗斯拉夫尔() 	yaroslavl.YaroslavlCounty
}

type 罗斯托夫RostovDuke struct {
	feud.BaseDuke
	别洛奥泽罗Beloozero 	beloozero.BeloozeroCounty
	罗斯托夫Rostov 	rostov.RostovCounty
	乌格利奇Uglich 	uglich.UglichCounty
	雅罗斯拉夫尔Yaroslavl 	yaroslavl.YaroslavlCounty
}

func (d *罗斯托夫RostovDuke) CBeloozero别洛奥泽罗() beloozero.BeloozeroCounty {
	return d.别洛奥泽罗Beloozero
}
    
func (d *罗斯托夫RostovDuke) CRostov罗斯托夫() rostov.RostovCounty {
	return d.罗斯托夫Rostov
}
    
func (d *罗斯托夫RostovDuke) CUglich乌格利奇() uglich.UglichCounty {
	return d.乌格利奇Uglich
}
    
func (d *罗斯托夫RostovDuke) CYaroslavl雅罗斯拉夫尔() yaroslavl.YaroslavlCounty {
	return d.雅罗斯拉夫尔Yaroslavl
}
    
var DRostov罗斯托夫 RostovDuke = &罗斯托夫RostovDuke{}

func init() {
	f := DRostov罗斯托夫.(*罗斯托夫RostovDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "rostov",
		TitleName: "罗斯托夫",
		TitleCode: "d_rostov",
		Counties:  map[string]feud.County{},
	}

	f.别洛奥泽罗Beloozero = beloozero.CBeloozero别洛奥泽罗
	f.别洛奥泽罗Beloozero.SetParent(f)
	
	f.罗斯托夫Rostov = rostov.CRostov罗斯托夫
	f.罗斯托夫Rostov.SetParent(f)
	
	f.乌格利奇Uglich = uglich.CUglich乌格利奇
	f.乌格利奇Uglich.SetParent(f)
	
	f.雅罗斯拉夫尔Yaroslavl = yaroslavl.CYaroslavl雅罗斯拉夫尔
	f.雅罗斯拉夫尔Yaroslavl.SetParent(f)
	
}
