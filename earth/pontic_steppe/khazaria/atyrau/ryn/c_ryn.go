package ryn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RynCounty interface {
    feud.County
    BBitik比蒂克() 	feud.Barony
    BBulan布兰() 	feud.Barony
    BChapaev恰帕耶沃() 	feud.Barony
    BOyan奥扬() 	feud.Barony
    BRyn雷恩() 	feud.Barony
    BShagatay沙加泰() 	feud.Barony
    BTendik坚季克() 	feud.Barony
}

type 雷恩RynCounty struct {
	feud.BaseCounty
	比蒂克Bitik 	feud.Barony
	布兰Bulan 	feud.Barony
	恰帕耶沃Chapaev 	feud.Barony
	奥扬Oyan 	feud.Barony
	雷恩Ryn 	feud.Barony
	沙加泰Shagatay 	feud.Barony
	坚季克Tendik 	feud.Barony
}

func (c *雷恩RynCounty) BBitik比蒂克() feud.Barony {
	return c.比蒂克Bitik
}
    
func (c *雷恩RynCounty) BBulan布兰() feud.Barony {
	return c.布兰Bulan
}
    
func (c *雷恩RynCounty) BChapaev恰帕耶沃() feud.Barony {
	return c.恰帕耶沃Chapaev
}
    
func (c *雷恩RynCounty) BOyan奥扬() feud.Barony {
	return c.奥扬Oyan
}
    
func (c *雷恩RynCounty) BRyn雷恩() feud.Barony {
	return c.雷恩Ryn
}
    
func (c *雷恩RynCounty) BShagatay沙加泰() feud.Barony {
	return c.沙加泰Shagatay
}
    
func (c *雷恩RynCounty) BTendik坚季克() feud.Barony {
	return c.坚季克Tendik
}
    
var CRyn雷恩 RynCounty = &雷恩RynCounty{}

func init() {
	f := CRyn雷恩.(*雷恩RynCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1813",
		Title:     "ryn",
		TitleName: "雷恩",
		TitleCode: "c_ryn",
		Baronies:  map[string]feud.Barony{},
	}

	f.比蒂克Bitik = BBitik比蒂克
	f.比蒂克Bitik.SetParent(f)
	
	f.布兰Bulan = BBulan布兰
	f.布兰Bulan.SetParent(f)
	
	f.恰帕耶沃Chapaev = BChapaev恰帕耶沃
	f.恰帕耶沃Chapaev.SetParent(f)
	
	f.奥扬Oyan = BOyan奥扬
	f.奥扬Oyan.SetParent(f)
	
	f.雷恩Ryn = BRyn雷恩
	f.雷恩Ryn.SetParent(f)
	
	f.沙加泰Shagatay = BShagatay沙加泰
	f.沙加泰Shagatay.SetParent(f)
	
	f.坚季克Tendik = BTendik坚季克
	f.坚季克Tendik.SetParent(f)
	
}
