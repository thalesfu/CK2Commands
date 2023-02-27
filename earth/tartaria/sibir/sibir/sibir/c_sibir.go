package sibir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SibirCounty interface {
    feud.County
    BBaduk巴杜克() 	feud.Barony
    BBelyyyar白亚尔() 	feud.Barony
    BIberbolgar伊别尔_保加尔() 	feud.Barony
    BKaik凯克() 	feud.Barony
    BLangepas兰格帕斯() 	feud.Barony
    BPokachi波卡奇() 	feud.Barony
    BSibir失必儿() 	feud.Barony
    BSurgut苏尔古特() 	feud.Barony
    BVysokiy维索基() 	feud.Barony
}

type 失必儿SibirCounty struct {
	feud.BaseCounty
	巴杜克Baduk 	feud.Barony
	白亚尔Belyyyar 	feud.Barony
	伊别尔_保加尔Iberbolgar 	feud.Barony
	凯克Kaik 	feud.Barony
	兰格帕斯Langepas 	feud.Barony
	波卡奇Pokachi 	feud.Barony
	失必儿Sibir 	feud.Barony
	苏尔古特Surgut 	feud.Barony
	维索基Vysokiy 	feud.Barony
}

func (c *失必儿SibirCounty) BBaduk巴杜克() feud.Barony {
	return c.巴杜克Baduk
}
    
func (c *失必儿SibirCounty) BBelyyyar白亚尔() feud.Barony {
	return c.白亚尔Belyyyar
}
    
func (c *失必儿SibirCounty) BIberbolgar伊别尔_保加尔() feud.Barony {
	return c.伊别尔_保加尔Iberbolgar
}
    
func (c *失必儿SibirCounty) BKaik凯克() feud.Barony {
	return c.凯克Kaik
}
    
func (c *失必儿SibirCounty) BLangepas兰格帕斯() feud.Barony {
	return c.兰格帕斯Langepas
}
    
func (c *失必儿SibirCounty) BPokachi波卡奇() feud.Barony {
	return c.波卡奇Pokachi
}
    
func (c *失必儿SibirCounty) BSibir失必儿() feud.Barony {
	return c.失必儿Sibir
}
    
func (c *失必儿SibirCounty) BSurgut苏尔古特() feud.Barony {
	return c.苏尔古特Surgut
}
    
func (c *失必儿SibirCounty) BVysokiy维索基() feud.Barony {
	return c.维索基Vysokiy
}
    
var CSibir失必儿 SibirCounty = &失必儿SibirCounty{}

func init() {
	f := CSibir失必儿.(*失必儿SibirCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "893",
		Title:     "sibir",
		TitleName: "失必儿",
		TitleCode: "c_sibir",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴杜克Baduk = BBaduk巴杜克
	f.巴杜克Baduk.SetParent(f)
	
	f.白亚尔Belyyyar = BBelyyyar白亚尔
	f.白亚尔Belyyyar.SetParent(f)
	
	f.伊别尔_保加尔Iberbolgar = BIberbolgar伊别尔_保加尔
	f.伊别尔_保加尔Iberbolgar.SetParent(f)
	
	f.凯克Kaik = BKaik凯克
	f.凯克Kaik.SetParent(f)
	
	f.兰格帕斯Langepas = BLangepas兰格帕斯
	f.兰格帕斯Langepas.SetParent(f)
	
	f.波卡奇Pokachi = BPokachi波卡奇
	f.波卡奇Pokachi.SetParent(f)
	
	f.失必儿Sibir = BSibir失必儿
	f.失必儿Sibir.SetParent(f)
	
	f.苏尔古特Surgut = BSurgut苏尔古特
	f.苏尔古特Surgut.SetParent(f)
	
	f.维索基Vysokiy = BVysokiy维索基
	f.维索基Vysokiy.SetParent(f)
	
}
