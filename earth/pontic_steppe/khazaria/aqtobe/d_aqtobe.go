package aqtobe

import (
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/aqtobe/aqtobe"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/aqtobe/inder"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/aqtobe/mughalzhar"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/aqtobe/utva"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AqtobeDuke interface {
    feud.Duke
    CAqtobe阿克托别() 	aqtobe.AqtobeCounty
    CInder伊列克() 	inder.InderCounty
    CMughalzhar穆加尔扎尔() 	mughalzhar.MughalzharCounty
    CUtva乌特瓦() 	utva.UtvaCounty
}

type 阿克托别AqtobeDuke struct {
	feud.BaseDuke
	阿克托别Aqtobe 	aqtobe.AqtobeCounty
	伊列克Inder 	inder.InderCounty
	穆加尔扎尔Mughalzhar 	mughalzhar.MughalzharCounty
	乌特瓦Utva 	utva.UtvaCounty
}

func (d *阿克托别AqtobeDuke) CAqtobe阿克托别() aqtobe.AqtobeCounty {
	return d.阿克托别Aqtobe
}
    
func (d *阿克托别AqtobeDuke) CInder伊列克() inder.InderCounty {
	return d.伊列克Inder
}
    
func (d *阿克托别AqtobeDuke) CMughalzhar穆加尔扎尔() mughalzhar.MughalzharCounty {
	return d.穆加尔扎尔Mughalzhar
}
    
func (d *阿克托别AqtobeDuke) CUtva乌特瓦() utva.UtvaCounty {
	return d.乌特瓦Utva
}
    
var DAqtobe阿克托别 AqtobeDuke = &阿克托别AqtobeDuke{}

func init() {
	f := DAqtobe阿克托别.(*阿克托别AqtobeDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "aqtobe",
		TitleName: "阿克托别",
		TitleCode: "d_aqtobe",
		Counties:  map[string]feud.County{},
	}

	f.阿克托别Aqtobe = aqtobe.CAqtobe阿克托别
	f.阿克托别Aqtobe.SetParent(f)
	
	f.伊列克Inder = inder.CInder伊列克
	f.伊列克Inder.SetParent(f)
	
	f.穆加尔扎尔Mughalzhar = mughalzhar.CMughalzhar穆加尔扎尔
	f.穆加尔扎尔Mughalzhar.SetParent(f)
	
	f.乌特瓦Utva = utva.CUtva乌特瓦
	f.乌特瓦Utva.SetParent(f)
	
}
