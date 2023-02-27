package vasyugan

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/vasyugan/bolshoy"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/vasyugan/shish"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/vasyugan/tui"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/vasyugan/vasyugan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VasyuganDuke interface {
    feud.Duke
    CBolshoy大尤甘() 	bolshoy.BolshoyCounty
    CShish希什() 	shish.ShishCounty
    CTui图伊() 	tui.TuiCounty
    CVasyugan瓦休甘() 	vasyugan.VasyuganCounty
}

type 瓦休甘VasyuganDuke struct {
	feud.BaseDuke
	大尤甘Bolshoy 	bolshoy.BolshoyCounty
	希什Shish 	shish.ShishCounty
	图伊Tui 	tui.TuiCounty
	瓦休甘Vasyugan 	vasyugan.VasyuganCounty
}

func (d *瓦休甘VasyuganDuke) CBolshoy大尤甘() bolshoy.BolshoyCounty {
	return d.大尤甘Bolshoy
}
    
func (d *瓦休甘VasyuganDuke) CShish希什() shish.ShishCounty {
	return d.希什Shish
}
    
func (d *瓦休甘VasyuganDuke) CTui图伊() tui.TuiCounty {
	return d.图伊Tui
}
    
func (d *瓦休甘VasyuganDuke) CVasyugan瓦休甘() vasyugan.VasyuganCounty {
	return d.瓦休甘Vasyugan
}
    
var DVasyugan瓦休甘 VasyuganDuke = &瓦休甘VasyuganDuke{}

func init() {
	f := DVasyugan瓦休甘.(*瓦休甘VasyuganDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "vasyugan",
		TitleName: "瓦休甘",
		TitleCode: "d_vasyugan",
		Counties:  map[string]feud.County{},
	}

	f.大尤甘Bolshoy = bolshoy.CBolshoy大尤甘
	f.大尤甘Bolshoy.SetParent(f)
	
	f.希什Shish = shish.CShish希什
	f.希什Shish.SetParent(f)
	
	f.图伊Tui = tui.CTui图伊
	f.图伊Tui.SetParent(f)
	
	f.瓦休甘Vasyugan = vasyugan.CVasyugan瓦休甘
	f.瓦休甘Vasyugan.SetParent(f)
	
}
