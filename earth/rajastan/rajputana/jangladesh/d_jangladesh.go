package jangladesh

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/jangladesh/nagauda"
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/jangladesh/reni"
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/jangladesh/vikramapura"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JangladeshDuke interface {
    feud.Duke
    CNagauda那乔佗() 	nagauda.NagaudaCounty
    CReni梨尼() 	reni.ReniCounty
    CVikramapura毗讫罗摩补罗() 	vikramapura.VikramapuraCounty
}

type 常伽罗提舍JangladeshDuke struct {
	feud.BaseDuke
	那乔佗Nagauda 	nagauda.NagaudaCounty
	梨尼Reni 	reni.ReniCounty
	毗讫罗摩补罗Vikramapura 	vikramapura.VikramapuraCounty
}

func (d *常伽罗提舍JangladeshDuke) CNagauda那乔佗() nagauda.NagaudaCounty {
	return d.那乔佗Nagauda
}
    
func (d *常伽罗提舍JangladeshDuke) CReni梨尼() reni.ReniCounty {
	return d.梨尼Reni
}
    
func (d *常伽罗提舍JangladeshDuke) CVikramapura毗讫罗摩补罗() vikramapura.VikramapuraCounty {
	return d.毗讫罗摩补罗Vikramapura
}
    
var DJangladesh常伽罗提舍 JangladeshDuke = &常伽罗提舍JangladeshDuke{}

func init() {
	f := DJangladesh常伽罗提舍.(*常伽罗提舍JangladeshDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "jangladesh",
		TitleName: "常伽罗提舍",
		TitleCode: "d_jangladesh",
		Counties:  map[string]feud.County{},
	}

	f.那乔佗Nagauda = nagauda.CNagauda那乔佗
	f.那乔佗Nagauda.SetParent(f)
	
	f.梨尼Reni = reni.CReni梨尼
	f.梨尼Reni.SetParent(f)
	
	f.毗讫罗摩补罗Vikramapura = vikramapura.CVikramapura毗讫罗摩补罗
	f.毗讫罗摩补罗Vikramapura.SetParent(f)
	
}
