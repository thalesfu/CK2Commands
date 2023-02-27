package czersk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CzerskCounty interface {
    feud.County
    BBrodno布鲁德诺() 	feud.Barony
    BLiw利夫() 	feud.Barony
    BLiw_nowy新利夫() 	feud.Barony
    BLochow沃胡夫() 	feud.Barony
    BRadzymin拉济明() 	feud.Barony
    BTluszoz特武什奇() 	feud.Barony
    BWegrow文格鲁夫() 	feud.Barony
}

type 利夫CzerskCounty struct {
	feud.BaseCounty
	布鲁德诺Brodno 	feud.Barony
	利夫Liw 	feud.Barony
	新利夫Liw_nowy 	feud.Barony
	沃胡夫Lochow 	feud.Barony
	拉济明Radzymin 	feud.Barony
	特武什奇Tluszoz 	feud.Barony
	文格鲁夫Wegrow 	feud.Barony
}

func (c *利夫CzerskCounty) BBrodno布鲁德诺() feud.Barony {
	return c.布鲁德诺Brodno
}
    
func (c *利夫CzerskCounty) BLiw利夫() feud.Barony {
	return c.利夫Liw
}
    
func (c *利夫CzerskCounty) BLiw_nowy新利夫() feud.Barony {
	return c.新利夫Liw_nowy
}
    
func (c *利夫CzerskCounty) BLochow沃胡夫() feud.Barony {
	return c.沃胡夫Lochow
}
    
func (c *利夫CzerskCounty) BRadzymin拉济明() feud.Barony {
	return c.拉济明Radzymin
}
    
func (c *利夫CzerskCounty) BTluszoz特武什奇() feud.Barony {
	return c.特武什奇Tluszoz
}
    
func (c *利夫CzerskCounty) BWegrow文格鲁夫() feud.Barony {
	return c.文格鲁夫Wegrow
}
    
var CCzersk利夫 CzerskCounty = &利夫CzerskCounty{}

func init() {
	f := CCzersk利夫.(*利夫CzerskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "530",
		Title:     "czersk",
		TitleName: "利夫",
		TitleCode: "c_czersk",
		Baronies:  map[string]feud.Barony{},
	}

	f.布鲁德诺Brodno = BBrodno布鲁德诺
	f.布鲁德诺Brodno.SetParent(f)
	
	f.利夫Liw = BLiw利夫
	f.利夫Liw.SetParent(f)
	
	f.新利夫Liw_nowy = BLiw_nowy新利夫
	f.新利夫Liw_nowy.SetParent(f)
	
	f.沃胡夫Lochow = BLochow沃胡夫
	f.沃胡夫Lochow.SetParent(f)
	
	f.拉济明Radzymin = BRadzymin拉济明
	f.拉济明Radzymin.SetParent(f)
	
	f.特武什奇Tluszoz = BTluszoz特武什奇
	f.特武什奇Tluszoz.SetParent(f)
	
	f.文格鲁夫Wegrow = BWegrow文格鲁夫
	f.文格鲁夫Wegrow.SetParent(f)
	
}
