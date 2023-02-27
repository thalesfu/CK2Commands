package olomouc

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OlomoucCounty interface {
    feud.County
    BHradiste城堡山() 	feud.Barony
    BOlomouc奥洛穆茨() 	feud.Barony
    BPrerov普热罗夫() 	feud.Barony
    BSvaty_kopecek圣科佩切克() 	feud.Barony
    BUnicov乌尼乔夫() 	feud.Barony
    BVelehrad韦莱赫拉德() 	feud.Barony
    BVyskov维什科夫() 	feud.Barony
    BZabreh扎布热赫() 	feud.Barony
}

type 奥洛穆茨OlomoucCounty struct {
	feud.BaseCounty
	城堡山Hradiste 	feud.Barony
	奥洛穆茨Olomouc 	feud.Barony
	普热罗夫Prerov 	feud.Barony
	圣科佩切克Svaty_kopecek 	feud.Barony
	乌尼乔夫Unicov 	feud.Barony
	韦莱赫拉德Velehrad 	feud.Barony
	维什科夫Vyskov 	feud.Barony
	扎布热赫Zabreh 	feud.Barony
}

func (c *奥洛穆茨OlomoucCounty) BHradiste城堡山() feud.Barony {
	return c.城堡山Hradiste
}
    
func (c *奥洛穆茨OlomoucCounty) BOlomouc奥洛穆茨() feud.Barony {
	return c.奥洛穆茨Olomouc
}
    
func (c *奥洛穆茨OlomoucCounty) BPrerov普热罗夫() feud.Barony {
	return c.普热罗夫Prerov
}
    
func (c *奥洛穆茨OlomoucCounty) BSvaty_kopecek圣科佩切克() feud.Barony {
	return c.圣科佩切克Svaty_kopecek
}
    
func (c *奥洛穆茨OlomoucCounty) BUnicov乌尼乔夫() feud.Barony {
	return c.乌尼乔夫Unicov
}
    
func (c *奥洛穆茨OlomoucCounty) BVelehrad韦莱赫拉德() feud.Barony {
	return c.韦莱赫拉德Velehrad
}
    
func (c *奥洛穆茨OlomoucCounty) BVyskov维什科夫() feud.Barony {
	return c.维什科夫Vyskov
}
    
func (c *奥洛穆茨OlomoucCounty) BZabreh扎布热赫() feud.Barony {
	return c.扎布热赫Zabreh
}
    
var COlomouc奥洛穆茨 OlomoucCounty = &奥洛穆茨OlomoucCounty{}

func init() {
	f := COlomouc奥洛穆茨.(*奥洛穆茨OlomoucCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "440",
		Title:     "olomouc",
		TitleName: "奥洛穆茨",
		TitleCode: "c_olomouc",
		Baronies:  map[string]feud.Barony{},
	}

	f.城堡山Hradiste = BHradiste城堡山
	f.城堡山Hradiste.SetParent(f)
	
	f.奥洛穆茨Olomouc = BOlomouc奥洛穆茨
	f.奥洛穆茨Olomouc.SetParent(f)
	
	f.普热罗夫Prerov = BPrerov普热罗夫
	f.普热罗夫Prerov.SetParent(f)
	
	f.圣科佩切克Svaty_kopecek = BSvaty_kopecek圣科佩切克
	f.圣科佩切克Svaty_kopecek.SetParent(f)
	
	f.乌尼乔夫Unicov = BUnicov乌尼乔夫
	f.乌尼乔夫Unicov.SetParent(f)
	
	f.韦莱赫拉德Velehrad = BVelehrad韦莱赫拉德
	f.韦莱赫拉德Velehrad.SetParent(f)
	
	f.维什科夫Vyskov = BVyskov维什科夫
	f.维什科夫Vyskov.SetParent(f)
	
	f.扎布热赫Zabreh = BZabreh扎布热赫
	f.扎布热赫Zabreh.SetParent(f)
	
}
