package rajanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RajanpurCounty interface {
	feud.County
	BAdhiwala代希瓦勒() feud.Barony
	BDajal陀迦() feud.Barony
	BHarrand哈兰德() feud.Barony
	BPottangi菩檀耆() feud.Barony
	BRajanpur罗旬补罗() feud.Barony
	BRojhan罗申() feud.Barony
	BSersandu瑟桑堵() feud.Barony
}

type 罗旬补罗RajanpurCounty struct {
	feud.BaseCounty
	代希瓦勒Adhiwala feud.Barony
	陀迦Dajal      feud.Barony
	哈兰德Harrand   feud.Barony
	菩檀耆Pottangi  feud.Barony
	罗旬补罗Rajanpur feud.Barony
	罗申Rojhan     feud.Barony
	瑟桑堵Sersandu  feud.Barony
}

func (c *罗旬补罗RajanpurCounty) BAdhiwala代希瓦勒() feud.Barony {
	return c.代希瓦勒Adhiwala
}

func (c *罗旬补罗RajanpurCounty) BDajal陀迦() feud.Barony {
	return c.陀迦Dajal
}

func (c *罗旬补罗RajanpurCounty) BHarrand哈兰德() feud.Barony {
	return c.哈兰德Harrand
}

func (c *罗旬补罗RajanpurCounty) BPottangi菩檀耆() feud.Barony {
	return c.菩檀耆Pottangi
}

func (c *罗旬补罗RajanpurCounty) BRajanpur罗旬补罗() feud.Barony {
	return c.罗旬补罗Rajanpur
}

func (c *罗旬补罗RajanpurCounty) BRojhan罗申() feud.Barony {
	return c.罗申Rojhan
}

func (c *罗旬补罗RajanpurCounty) BSersandu瑟桑堵() feud.Barony {
	return c.瑟桑堵Sersandu
}

var CRajanpur罗旬补罗 RajanpurCounty = &罗旬补罗RajanpurCounty{}

func init() {
	f := CRajanpur罗旬补罗.(*罗旬补罗RajanpurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1339",
		Title:     "rajanpur",
		TitleName: "罗旬补罗",
		TitleCode: "c_rajanpur",
		Baronies:  map[string]feud.Barony{},
	}

	f.代希瓦勒Adhiwala = BAdhiwala代希瓦勒
	f.代希瓦勒Adhiwala.SetParent(f)

	f.陀迦Dajal = BDajal陀迦
	f.陀迦Dajal.SetParent(f)

	f.哈兰德Harrand = BHarrand哈兰德
	f.哈兰德Harrand.SetParent(f)

	f.菩檀耆Pottangi = BPottangi菩檀耆
	f.菩檀耆Pottangi.SetParent(f)

	f.罗旬补罗Rajanpur = BRajanpur罗旬补罗
	f.罗旬补罗Rajanpur.SetParent(f)

	f.罗申Rojhan = BRojhan罗申
	f.罗申Rojhan.SetParent(f)

	f.瑟桑堵Sersandu = BSersandu瑟桑堵
	f.瑟桑堵Sersandu.SetParent(f)

}
