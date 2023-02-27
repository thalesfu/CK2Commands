package salamanca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SalamancaCounty interface {
    feud.County
    BAlbadetormes托梅斯河畔阿尔瓦() 	feud.Barony
    BBracamonte布拉卡蒙特() 	feud.Barony
    BCarbajosadelasagrada卡尔瓦霍萨德拉萨格拉达() 	feud.Barony
    BCiudadrodrigo罗德里戈城() 	feud.Barony
    BLumbrales伦布拉莱斯() 	feud.Barony
    BSalamanca萨拉曼卡() 	feud.Barony
    BSalbejar贝哈尔() 	feud.Barony
    BTerradillos特拉迪略斯() 	feud.Barony
}

type 萨拉曼卡SalamancaCounty struct {
	feud.BaseCounty
	托梅斯河畔阿尔瓦Albadetormes 	feud.Barony
	布拉卡蒙特Bracamonte 	feud.Barony
	卡尔瓦霍萨德拉萨格拉达Carbajosadelasagrada 	feud.Barony
	罗德里戈城Ciudadrodrigo 	feud.Barony
	伦布拉莱斯Lumbrales 	feud.Barony
	萨拉曼卡Salamanca 	feud.Barony
	贝哈尔Salbejar 	feud.Barony
	特拉迪略斯Terradillos 	feud.Barony
}

func (c *萨拉曼卡SalamancaCounty) BAlbadetormes托梅斯河畔阿尔瓦() feud.Barony {
	return c.托梅斯河畔阿尔瓦Albadetormes
}
    
func (c *萨拉曼卡SalamancaCounty) BBracamonte布拉卡蒙特() feud.Barony {
	return c.布拉卡蒙特Bracamonte
}
    
func (c *萨拉曼卡SalamancaCounty) BCarbajosadelasagrada卡尔瓦霍萨德拉萨格拉达() feud.Barony {
	return c.卡尔瓦霍萨德拉萨格拉达Carbajosadelasagrada
}
    
func (c *萨拉曼卡SalamancaCounty) BCiudadrodrigo罗德里戈城() feud.Barony {
	return c.罗德里戈城Ciudadrodrigo
}
    
func (c *萨拉曼卡SalamancaCounty) BLumbrales伦布拉莱斯() feud.Barony {
	return c.伦布拉莱斯Lumbrales
}
    
func (c *萨拉曼卡SalamancaCounty) BSalamanca萨拉曼卡() feud.Barony {
	return c.萨拉曼卡Salamanca
}
    
func (c *萨拉曼卡SalamancaCounty) BSalbejar贝哈尔() feud.Barony {
	return c.贝哈尔Salbejar
}
    
func (c *萨拉曼卡SalamancaCounty) BTerradillos特拉迪略斯() feud.Barony {
	return c.特拉迪略斯Terradillos
}
    
var CSalamanca萨拉曼卡 SalamancaCounty = &萨拉曼卡SalamancaCounty{}

func init() {
	f := CSalamanca萨拉曼卡.(*萨拉曼卡SalamancaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "192",
		Title:     "salamanca",
		TitleName: "萨拉曼卡",
		TitleCode: "c_salamanca",
		Baronies:  map[string]feud.Barony{},
	}

	f.托梅斯河畔阿尔瓦Albadetormes = BAlbadetormes托梅斯河畔阿尔瓦
	f.托梅斯河畔阿尔瓦Albadetormes.SetParent(f)
	
	f.布拉卡蒙特Bracamonte = BBracamonte布拉卡蒙特
	f.布拉卡蒙特Bracamonte.SetParent(f)
	
	f.卡尔瓦霍萨德拉萨格拉达Carbajosadelasagrada = BCarbajosadelasagrada卡尔瓦霍萨德拉萨格拉达
	f.卡尔瓦霍萨德拉萨格拉达Carbajosadelasagrada.SetParent(f)
	
	f.罗德里戈城Ciudadrodrigo = BCiudadrodrigo罗德里戈城
	f.罗德里戈城Ciudadrodrigo.SetParent(f)
	
	f.伦布拉莱斯Lumbrales = BLumbrales伦布拉莱斯
	f.伦布拉莱斯Lumbrales.SetParent(f)
	
	f.萨拉曼卡Salamanca = BSalamanca萨拉曼卡
	f.萨拉曼卡Salamanca.SetParent(f)
	
	f.贝哈尔Salbejar = BSalbejar贝哈尔
	f.贝哈尔Salbejar.SetParent(f)
	
	f.特拉迪略斯Terradillos = BTerradillos特拉迪略斯
	f.特拉迪略斯Terradillos.SetParent(f)
	
}
