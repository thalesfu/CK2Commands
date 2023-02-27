package vladimir

import (
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/vladimir/gorodez"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/vladimir/nizhny_novgorod"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/vladimir/pereyaslavl_zalessky"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/vladimir/suzdal"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/vladimir/vladimir"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/vladimir/yuryev"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VladimirDuke interface {
    feud.Duke
    CGorodez戈罗杰茨() 	gorodez.GorodezCounty
    CNizhny_novgorod下诺夫哥罗德() 	nizhny_novgorod.Nizhny_novgorodCounty
    CPereyaslavl_zalessky佩列亚斯拉夫尔扎列斯基() 	pereyaslavl_zalessky.Pereyaslavl_zalesskyCounty
    CSuzdal苏兹达尔() 	suzdal.SuzdalCounty
    CVladimir弗拉基米尔() 	vladimir.VladimirCounty
    CYuryev尤里耶夫() 	yuryev.YuryevCounty
}

type 弗拉基米尔VladimirDuke struct {
	feud.BaseDuke
	戈罗杰茨Gorodez 	gorodez.GorodezCounty
	下诺夫哥罗德Nizhny_novgorod 	nizhny_novgorod.Nizhny_novgorodCounty
	佩列亚斯拉夫尔扎列斯基Pereyaslavl_zalessky 	pereyaslavl_zalessky.Pereyaslavl_zalesskyCounty
	苏兹达尔Suzdal 	suzdal.SuzdalCounty
	弗拉基米尔Vladimir 	vladimir.VladimirCounty
	尤里耶夫Yuryev 	yuryev.YuryevCounty
}

func (d *弗拉基米尔VladimirDuke) CGorodez戈罗杰茨() gorodez.GorodezCounty {
	return d.戈罗杰茨Gorodez
}
    
func (d *弗拉基米尔VladimirDuke) CNizhny_novgorod下诺夫哥罗德() nizhny_novgorod.Nizhny_novgorodCounty {
	return d.下诺夫哥罗德Nizhny_novgorod
}
    
func (d *弗拉基米尔VladimirDuke) CPereyaslavl_zalessky佩列亚斯拉夫尔扎列斯基() pereyaslavl_zalessky.Pereyaslavl_zalesskyCounty {
	return d.佩列亚斯拉夫尔扎列斯基Pereyaslavl_zalessky
}
    
func (d *弗拉基米尔VladimirDuke) CSuzdal苏兹达尔() suzdal.SuzdalCounty {
	return d.苏兹达尔Suzdal
}
    
func (d *弗拉基米尔VladimirDuke) CVladimir弗拉基米尔() vladimir.VladimirCounty {
	return d.弗拉基米尔Vladimir
}
    
func (d *弗拉基米尔VladimirDuke) CYuryev尤里耶夫() yuryev.YuryevCounty {
	return d.尤里耶夫Yuryev
}
    
var DVladimir弗拉基米尔 VladimirDuke = &弗拉基米尔VladimirDuke{}

func init() {
	f := DVladimir弗拉基米尔.(*弗拉基米尔VladimirDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "vladimir",
		TitleName: "弗拉基米尔",
		TitleCode: "d_vladimir",
		Counties:  map[string]feud.County{},
	}

	f.戈罗杰茨Gorodez = gorodez.CGorodez戈罗杰茨
	f.戈罗杰茨Gorodez.SetParent(f)
	
	f.下诺夫哥罗德Nizhny_novgorod = nizhny_novgorod.CNizhny_novgorod下诺夫哥罗德
	f.下诺夫哥罗德Nizhny_novgorod.SetParent(f)
	
	f.佩列亚斯拉夫尔扎列斯基Pereyaslavl_zalessky = pereyaslavl_zalessky.CPereyaslavl_zalessky佩列亚斯拉夫尔扎列斯基
	f.佩列亚斯拉夫尔扎列斯基Pereyaslavl_zalessky.SetParent(f)
	
	f.苏兹达尔Suzdal = suzdal.CSuzdal苏兹达尔
	f.苏兹达尔Suzdal.SetParent(f)
	
	f.弗拉基米尔Vladimir = vladimir.CVladimir弗拉基米尔
	f.弗拉基米尔Vladimir.SetParent(f)
	
	f.尤里耶夫Yuryev = yuryev.CYuryev尤里耶夫
	f.尤里耶夫Yuryev.SetParent(f)
	
}
