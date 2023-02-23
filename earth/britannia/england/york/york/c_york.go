package york

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YorkCounty interface {
	feud.County
	BMiddlesbrough米德尔斯堡() feud.Barony
	BRichmond里奇蒙() feud.Barony
	BScarborough斯卡布罗() feud.Barony
	BThirsk瑟斯克() feud.Barony
	BWhitby惠特比() feud.Barony
	BYork约克() feud.Barony
}

type 约克YorkCounty struct {
	feud.BaseCounty
	米德尔斯堡Middlesbrough feud.Barony
	里奇蒙Richmond        feud.Barony
	斯卡布罗Scarborough    feud.Barony
	瑟斯克Thirsk          feud.Barony
	惠特比Whitby          feud.Barony
	约克York             feud.Barony
}

func (c *约克YorkCounty) BMiddlesbrough米德尔斯堡() feud.Barony {
	return c.米德尔斯堡Middlesbrough
}

func (c *约克YorkCounty) BRichmond里奇蒙() feud.Barony {
	return c.里奇蒙Richmond
}

func (c *约克YorkCounty) BScarborough斯卡布罗() feud.Barony {
	return c.斯卡布罗Scarborough
}

func (c *约克YorkCounty) BThirsk瑟斯克() feud.Barony {
	return c.瑟斯克Thirsk
}

func (c *约克YorkCounty) BWhitby惠特比() feud.Barony {
	return c.惠特比Whitby
}

func (c *约克YorkCounty) BYork约克() feud.Barony {
	return c.约克York
}

var CYork约克 YorkCounty = &约克YorkCounty{}

func init() {
	f := CYork约克.(*约克YorkCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "57",
		Title:     "york",
		TitleName: "约克",
		TitleCode: "c_york",
		Baronies:  map[string]feud.Barony{},
	}

	f.米德尔斯堡Middlesbrough = BMiddlesbrough米德尔斯堡
	f.米德尔斯堡Middlesbrough.SetParent(f)

	f.里奇蒙Richmond = BRichmond里奇蒙
	f.里奇蒙Richmond.SetParent(f)

	f.斯卡布罗Scarborough = BScarborough斯卡布罗
	f.斯卡布罗Scarborough.SetParent(f)

	f.瑟斯克Thirsk = BThirsk瑟斯克
	f.瑟斯克Thirsk.SetParent(f)

	f.惠特比Whitby = BWhitby惠特比
	f.惠特比Whitby.SetParent(f)

	f.约克York = BYork约克
	f.约克York.SetParent(f)

}
