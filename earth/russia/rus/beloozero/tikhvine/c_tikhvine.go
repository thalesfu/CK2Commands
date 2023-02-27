package tikhvine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TikhvineCounty interface {
    feud.County
    BBabayevo巴巴耶沃() 	feud.Barony
    BBoksiti博克西特() 	feud.Barony
    BChagoda恰戈达() 	feud.Barony
    BKhvoynaya赫沃伊纳亚() 	feud.Barony
    BPestovo佩斯托沃() 	feud.Barony
    BPikalyovo皮卡廖沃() 	feud.Barony
    BTikhvine季赫温() 	feud.Barony
}

type 季赫温TikhvineCounty struct {
	feud.BaseCounty
	巴巴耶沃Babayevo 	feud.Barony
	博克西特Boksiti 	feud.Barony
	恰戈达Chagoda 	feud.Barony
	赫沃伊纳亚Khvoynaya 	feud.Barony
	佩斯托沃Pestovo 	feud.Barony
	皮卡廖沃Pikalyovo 	feud.Barony
	季赫温Tikhvine 	feud.Barony
}

func (c *季赫温TikhvineCounty) BBabayevo巴巴耶沃() feud.Barony {
	return c.巴巴耶沃Babayevo
}
    
func (c *季赫温TikhvineCounty) BBoksiti博克西特() feud.Barony {
	return c.博克西特Boksiti
}
    
func (c *季赫温TikhvineCounty) BChagoda恰戈达() feud.Barony {
	return c.恰戈达Chagoda
}
    
func (c *季赫温TikhvineCounty) BKhvoynaya赫沃伊纳亚() feud.Barony {
	return c.赫沃伊纳亚Khvoynaya
}
    
func (c *季赫温TikhvineCounty) BPestovo佩斯托沃() feud.Barony {
	return c.佩斯托沃Pestovo
}
    
func (c *季赫温TikhvineCounty) BPikalyovo皮卡廖沃() feud.Barony {
	return c.皮卡廖沃Pikalyovo
}
    
func (c *季赫温TikhvineCounty) BTikhvine季赫温() feud.Barony {
	return c.季赫温Tikhvine
}
    
var CTikhvine季赫温 TikhvineCounty = &季赫温TikhvineCounty{}

func init() {
	f := CTikhvine季赫温.(*季赫温TikhvineCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1664",
		Title:     "tikhvine",
		TitleName: "季赫温",
		TitleCode: "c_tikhvine",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴巴耶沃Babayevo = BBabayevo巴巴耶沃
	f.巴巴耶沃Babayevo.SetParent(f)
	
	f.博克西特Boksiti = BBoksiti博克西特
	f.博克西特Boksiti.SetParent(f)
	
	f.恰戈达Chagoda = BChagoda恰戈达
	f.恰戈达Chagoda.SetParent(f)
	
	f.赫沃伊纳亚Khvoynaya = BKhvoynaya赫沃伊纳亚
	f.赫沃伊纳亚Khvoynaya.SetParent(f)
	
	f.佩斯托沃Pestovo = BPestovo佩斯托沃
	f.佩斯托沃Pestovo.SetParent(f)
	
	f.皮卡廖沃Pikalyovo = BPikalyovo皮卡廖沃
	f.皮卡廖沃Pikalyovo.SetParent(f)
	
	f.季赫温Tikhvine = BTikhvine季赫温
	f.季赫温Tikhvine.SetParent(f)
	
}
