package mecklemburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MecklemburgCounty interface {
    feud.County
    BGevezin盖沃钦() 	feud.Barony
    BGrevesmuhlen格雷沃斯米伦() 	feud.Barony
    BHagenow哈格诺() 	feud.Barony
    BLutzow吕措() 	feud.Barony
    BMecklemburg梅克伦堡() 	feud.Barony
    BParchim帕尔希姆() 	feud.Barony
    BSchwerin什未林() 	feud.Barony
    BWismar维斯马() 	feud.Barony
}

type 梅克伦堡MecklemburgCounty struct {
	feud.BaseCounty
	盖沃钦Gevezin 	feud.Barony
	格雷沃斯米伦Grevesmuhlen 	feud.Barony
	哈格诺Hagenow 	feud.Barony
	吕措Lutzow 	feud.Barony
	梅克伦堡Mecklemburg 	feud.Barony
	帕尔希姆Parchim 	feud.Barony
	什未林Schwerin 	feud.Barony
	维斯马Wismar 	feud.Barony
}

func (c *梅克伦堡MecklemburgCounty) BGevezin盖沃钦() feud.Barony {
	return c.盖沃钦Gevezin
}
    
func (c *梅克伦堡MecklemburgCounty) BGrevesmuhlen格雷沃斯米伦() feud.Barony {
	return c.格雷沃斯米伦Grevesmuhlen
}
    
func (c *梅克伦堡MecklemburgCounty) BHagenow哈格诺() feud.Barony {
	return c.哈格诺Hagenow
}
    
func (c *梅克伦堡MecklemburgCounty) BLutzow吕措() feud.Barony {
	return c.吕措Lutzow
}
    
func (c *梅克伦堡MecklemburgCounty) BMecklemburg梅克伦堡() feud.Barony {
	return c.梅克伦堡Mecklemburg
}
    
func (c *梅克伦堡MecklemburgCounty) BParchim帕尔希姆() feud.Barony {
	return c.帕尔希姆Parchim
}
    
func (c *梅克伦堡MecklemburgCounty) BSchwerin什未林() feud.Barony {
	return c.什未林Schwerin
}
    
func (c *梅克伦堡MecklemburgCounty) BWismar维斯马() feud.Barony {
	return c.维斯马Wismar
}
    
var CMecklemburg梅克伦堡 MecklemburgCounty = &梅克伦堡MecklemburgCounty{}

func init() {
	f := CMecklemburg梅克伦堡.(*梅克伦堡MecklemburgCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "260",
		Title:     "mecklemburg",
		TitleName: "梅克伦堡",
		TitleCode: "c_mecklemburg",
		Baronies:  map[string]feud.Barony{},
	}

	f.盖沃钦Gevezin = BGevezin盖沃钦
	f.盖沃钦Gevezin.SetParent(f)
	
	f.格雷沃斯米伦Grevesmuhlen = BGrevesmuhlen格雷沃斯米伦
	f.格雷沃斯米伦Grevesmuhlen.SetParent(f)
	
	f.哈格诺Hagenow = BHagenow哈格诺
	f.哈格诺Hagenow.SetParent(f)
	
	f.吕措Lutzow = BLutzow吕措
	f.吕措Lutzow.SetParent(f)
	
	f.梅克伦堡Mecklemburg = BMecklemburg梅克伦堡
	f.梅克伦堡Mecklemburg.SetParent(f)
	
	f.帕尔希姆Parchim = BParchim帕尔希姆
	f.帕尔希姆Parchim.SetParent(f)
	
	f.什未林Schwerin = BSchwerin什未林
	f.什未林Schwerin.SetParent(f)
	
	f.维斯马Wismar = BWismar维斯马
	f.维斯马Wismar.SetParent(f)
	
}
