package blois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BloisCounty interface {
    feud.County
    BBeaugency博让西() 	feud.Barony
    BBlois布卢瓦() 	feud.Barony
    BChaumontsurloire卢瓦尔河畔肖蒙() 	feud.Barony
    BFougeressurbievre比耶夫尔河畔富热尔() 	feud.Barony
    BMontrichard蒙特里夏尔() 	feud.Barony
    BRomorantin罗莫朗坦() 	feud.Barony
    BStaignan圣艾尼昂() 	feud.Barony
    BStgeorgesdubois圣乔治迪布瓦() 	feud.Barony
}

type 布卢瓦BloisCounty struct {
	feud.BaseCounty
	博让西Beaugency 	feud.Barony
	布卢瓦Blois 	feud.Barony
	卢瓦尔河畔肖蒙Chaumontsurloire 	feud.Barony
	比耶夫尔河畔富热尔Fougeressurbievre 	feud.Barony
	蒙特里夏尔Montrichard 	feud.Barony
	罗莫朗坦Romorantin 	feud.Barony
	圣艾尼昂Staignan 	feud.Barony
	圣乔治迪布瓦Stgeorgesdubois 	feud.Barony
}

func (c *布卢瓦BloisCounty) BBeaugency博让西() feud.Barony {
	return c.博让西Beaugency
}
    
func (c *布卢瓦BloisCounty) BBlois布卢瓦() feud.Barony {
	return c.布卢瓦Blois
}
    
func (c *布卢瓦BloisCounty) BChaumontsurloire卢瓦尔河畔肖蒙() feud.Barony {
	return c.卢瓦尔河畔肖蒙Chaumontsurloire
}
    
func (c *布卢瓦BloisCounty) BFougeressurbievre比耶夫尔河畔富热尔() feud.Barony {
	return c.比耶夫尔河畔富热尔Fougeressurbievre
}
    
func (c *布卢瓦BloisCounty) BMontrichard蒙特里夏尔() feud.Barony {
	return c.蒙特里夏尔Montrichard
}
    
func (c *布卢瓦BloisCounty) BRomorantin罗莫朗坦() feud.Barony {
	return c.罗莫朗坦Romorantin
}
    
func (c *布卢瓦BloisCounty) BStaignan圣艾尼昂() feud.Barony {
	return c.圣艾尼昂Staignan
}
    
func (c *布卢瓦BloisCounty) BStgeorgesdubois圣乔治迪布瓦() feud.Barony {
	return c.圣乔治迪布瓦Stgeorgesdubois
}
    
var CBlois布卢瓦 BloisCounty = &布卢瓦BloisCounty{}

func init() {
	f := CBlois布卢瓦.(*布卢瓦BloisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "110",
		Title:     "blois",
		TitleName: "布卢瓦",
		TitleCode: "c_blois",
		Baronies:  map[string]feud.Barony{},
	}

	f.博让西Beaugency = BBeaugency博让西
	f.博让西Beaugency.SetParent(f)
	
	f.布卢瓦Blois = BBlois布卢瓦
	f.布卢瓦Blois.SetParent(f)
	
	f.卢瓦尔河畔肖蒙Chaumontsurloire = BChaumontsurloire卢瓦尔河畔肖蒙
	f.卢瓦尔河畔肖蒙Chaumontsurloire.SetParent(f)
	
	f.比耶夫尔河畔富热尔Fougeressurbievre = BFougeressurbievre比耶夫尔河畔富热尔
	f.比耶夫尔河畔富热尔Fougeressurbievre.SetParent(f)
	
	f.蒙特里夏尔Montrichard = BMontrichard蒙特里夏尔
	f.蒙特里夏尔Montrichard.SetParent(f)
	
	f.罗莫朗坦Romorantin = BRomorantin罗莫朗坦
	f.罗莫朗坦Romorantin.SetParent(f)
	
	f.圣艾尼昂Staignan = BStaignan圣艾尼昂
	f.圣艾尼昂Staignan.SetParent(f)
	
	f.圣乔治迪布瓦Stgeorgesdubois = BStgeorgesdubois圣乔治迪布瓦
	f.圣乔治迪布瓦Stgeorgesdubois.SetParent(f)
	
}
