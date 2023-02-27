package vijayapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VijayapuraCounty interface {
    feud.County
    BBegunia毗古尼阿() 	feud.Barony
    BChurulia周楼梨耶() 	feud.Barony
    BDishergarh提室罗姞利呬() 	feud.Barony
    BKalna迦罗那() 	feud.Barony
    BMohammadnagar摩诃末那揭罗() 	feud.Barony
    BPatuli波堵梨() 	feud.Barony
    BVijayapura毗阇耶城() 	feud.Barony
}

type 毗阇耶城VijayapuraCounty struct {
	feud.BaseCounty
	毗古尼阿Begunia 	feud.Barony
	周楼梨耶Churulia 	feud.Barony
	提室罗姞利呬Dishergarh 	feud.Barony
	迦罗那Kalna 	feud.Barony
	摩诃末那揭罗Mohammadnagar 	feud.Barony
	波堵梨Patuli 	feud.Barony
	毗阇耶城Vijayapura 	feud.Barony
}

func (c *毗阇耶城VijayapuraCounty) BBegunia毗古尼阿() feud.Barony {
	return c.毗古尼阿Begunia
}
    
func (c *毗阇耶城VijayapuraCounty) BChurulia周楼梨耶() feud.Barony {
	return c.周楼梨耶Churulia
}
    
func (c *毗阇耶城VijayapuraCounty) BDishergarh提室罗姞利呬() feud.Barony {
	return c.提室罗姞利呬Dishergarh
}
    
func (c *毗阇耶城VijayapuraCounty) BKalna迦罗那() feud.Barony {
	return c.迦罗那Kalna
}
    
func (c *毗阇耶城VijayapuraCounty) BMohammadnagar摩诃末那揭罗() feud.Barony {
	return c.摩诃末那揭罗Mohammadnagar
}
    
func (c *毗阇耶城VijayapuraCounty) BPatuli波堵梨() feud.Barony {
	return c.波堵梨Patuli
}
    
func (c *毗阇耶城VijayapuraCounty) BVijayapura毗阇耶城() feud.Barony {
	return c.毗阇耶城Vijayapura
}
    
var CVijayapura毗阇耶城 VijayapuraCounty = &毗阇耶城VijayapuraCounty{}

func init() {
	f := CVijayapura毗阇耶城.(*毗阇耶城VijayapuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1239",
		Title:     "vijayapura",
		TitleName: "毗阇耶城",
		TitleCode: "c_vijayapura",
		Baronies:  map[string]feud.Barony{},
	}

	f.毗古尼阿Begunia = BBegunia毗古尼阿
	f.毗古尼阿Begunia.SetParent(f)
	
	f.周楼梨耶Churulia = BChurulia周楼梨耶
	f.周楼梨耶Churulia.SetParent(f)
	
	f.提室罗姞利呬Dishergarh = BDishergarh提室罗姞利呬
	f.提室罗姞利呬Dishergarh.SetParent(f)
	
	f.迦罗那Kalna = BKalna迦罗那
	f.迦罗那Kalna.SetParent(f)
	
	f.摩诃末那揭罗Mohammadnagar = BMohammadnagar摩诃末那揭罗
	f.摩诃末那揭罗Mohammadnagar.SetParent(f)
	
	f.波堵梨Patuli = BPatuli波堵梨
	f.波堵梨Patuli.SetParent(f)
	
	f.毗阇耶城Vijayapura = BVijayapura毗阇耶城
	f.毗阇耶城Vijayapura.SetParent(f)
	
}
