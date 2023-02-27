package udine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UdineCounty interface {
    feud.County
    BCiridale奇维达莱() 	feud.Barony
    BFriuli塔尔琴托() 	feud.Barony
    BGemona杰莫纳() 	feud.Barony
    BManiago马尼亚戈() 	feud.Barony
    BTarcento塔尔琴托() 	feud.Barony
    BTolmezzo托尔梅佐() 	feud.Barony
    BUdine乌迪内() 	feud.Barony
}

type 乌迪内UdineCounty struct {
	feud.BaseCounty
	奇维达莱Ciridale 	feud.Barony
	塔尔琴托Friuli 	feud.Barony
	杰莫纳Gemona 	feud.Barony
	马尼亚戈Maniago 	feud.Barony
	塔尔琴托Tarcento 	feud.Barony
	托尔梅佐Tolmezzo 	feud.Barony
	乌迪内Udine 	feud.Barony
}

func (c *乌迪内UdineCounty) BCiridale奇维达莱() feud.Barony {
	return c.奇维达莱Ciridale
}
    
func (c *乌迪内UdineCounty) BFriuli塔尔琴托() feud.Barony {
	return c.塔尔琴托Friuli
}
    
func (c *乌迪内UdineCounty) BGemona杰莫纳() feud.Barony {
	return c.杰莫纳Gemona
}
    
func (c *乌迪内UdineCounty) BManiago马尼亚戈() feud.Barony {
	return c.马尼亚戈Maniago
}
    
func (c *乌迪内UdineCounty) BTarcento塔尔琴托() feud.Barony {
	return c.塔尔琴托Tarcento
}
    
func (c *乌迪内UdineCounty) BTolmezzo托尔梅佐() feud.Barony {
	return c.托尔梅佐Tolmezzo
}
    
func (c *乌迪内UdineCounty) BUdine乌迪内() feud.Barony {
	return c.乌迪内Udine
}
    
var CUdine乌迪内 UdineCounty = &乌迪内UdineCounty{}

func init() {
	f := CUdine乌迪内.(*乌迪内UdineCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1695",
		Title:     "udine",
		TitleName: "乌迪内",
		TitleCode: "c_udine",
		Baronies:  map[string]feud.Barony{},
	}

	f.奇维达莱Ciridale = BCiridale奇维达莱
	f.奇维达莱Ciridale.SetParent(f)
	
	f.塔尔琴托Friuli = BFriuli塔尔琴托
	f.塔尔琴托Friuli.SetParent(f)
	
	f.杰莫纳Gemona = BGemona杰莫纳
	f.杰莫纳Gemona.SetParent(f)
	
	f.马尼亚戈Maniago = BManiago马尼亚戈
	f.马尼亚戈Maniago.SetParent(f)
	
	f.塔尔琴托Tarcento = BTarcento塔尔琴托
	f.塔尔琴托Tarcento.SetParent(f)
	
	f.托尔梅佐Tolmezzo = BTolmezzo托尔梅佐
	f.托尔梅佐Tolmezzo.SetParent(f)
	
	f.乌迪内Udine = BUdine乌迪内
	f.乌迪内Udine.SetParent(f)
	
}
