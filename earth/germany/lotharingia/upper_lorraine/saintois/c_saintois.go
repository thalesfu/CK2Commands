package saintois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SaintoisCounty interface {
    feud.County
    BBrixey布里克塞() 	feud.Barony
    BDompaire栋派尔() 	feud.Barony
    BIvios伊维() 	feud.Barony
    BOrnois奥尔努瓦() 	feud.Barony
    BSaintois桑图瓦() 	feud.Barony
    BSaulnois索尔努瓦() 	feud.Barony
    BSorcystmartin索尔西圣马尔坦() 	feud.Barony
    BVaudemont沃代蒙() 	feud.Barony
}

type 桑图瓦SaintoisCounty struct {
	feud.BaseCounty
	布里克塞Brixey 	feud.Barony
	栋派尔Dompaire 	feud.Barony
	伊维Ivios 	feud.Barony
	奥尔努瓦Ornois 	feud.Barony
	桑图瓦Saintois 	feud.Barony
	索尔努瓦Saulnois 	feud.Barony
	索尔西圣马尔坦Sorcystmartin 	feud.Barony
	沃代蒙Vaudemont 	feud.Barony
}

func (c *桑图瓦SaintoisCounty) BBrixey布里克塞() feud.Barony {
	return c.布里克塞Brixey
}
    
func (c *桑图瓦SaintoisCounty) BDompaire栋派尔() feud.Barony {
	return c.栋派尔Dompaire
}
    
func (c *桑图瓦SaintoisCounty) BIvios伊维() feud.Barony {
	return c.伊维Ivios
}
    
func (c *桑图瓦SaintoisCounty) BOrnois奥尔努瓦() feud.Barony {
	return c.奥尔努瓦Ornois
}
    
func (c *桑图瓦SaintoisCounty) BSaintois桑图瓦() feud.Barony {
	return c.桑图瓦Saintois
}
    
func (c *桑图瓦SaintoisCounty) BSaulnois索尔努瓦() feud.Barony {
	return c.索尔努瓦Saulnois
}
    
func (c *桑图瓦SaintoisCounty) BSorcystmartin索尔西圣马尔坦() feud.Barony {
	return c.索尔西圣马尔坦Sorcystmartin
}
    
func (c *桑图瓦SaintoisCounty) BVaudemont沃代蒙() feud.Barony {
	return c.沃代蒙Vaudemont
}
    
var CSaintois桑图瓦 SaintoisCounty = &桑图瓦SaintoisCounty{}

func init() {
	f := CSaintois桑图瓦.(*桑图瓦SaintoisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "133",
		Title:     "saintois",
		TitleName: "桑图瓦",
		TitleCode: "c_saintois",
		Baronies:  map[string]feud.Barony{},
	}

	f.布里克塞Brixey = BBrixey布里克塞
	f.布里克塞Brixey.SetParent(f)
	
	f.栋派尔Dompaire = BDompaire栋派尔
	f.栋派尔Dompaire.SetParent(f)
	
	f.伊维Ivios = BIvios伊维
	f.伊维Ivios.SetParent(f)
	
	f.奥尔努瓦Ornois = BOrnois奥尔努瓦
	f.奥尔努瓦Ornois.SetParent(f)
	
	f.桑图瓦Saintois = BSaintois桑图瓦
	f.桑图瓦Saintois.SetParent(f)
	
	f.索尔努瓦Saulnois = BSaulnois索尔努瓦
	f.索尔努瓦Saulnois.SetParent(f)
	
	f.索尔西圣马尔坦Sorcystmartin = BSorcystmartin索尔西圣马尔坦
	f.索尔西圣马尔坦Sorcystmartin.SetParent(f)
	
	f.沃代蒙Vaudemont = BVaudemont沃代蒙
	f.沃代蒙Vaudemont.SetParent(f)
	
}
