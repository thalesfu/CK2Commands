package gurgi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GurgiCounty interface {
	feud.County
	BBathgahora婆多迦呼罗() feud.Barony
	BChandrehi旃陀梨醯() feud.Barony
	BGurgi古卢耆() feud.Barony
	BSathon沙吞() feud.Barony
	BSatna萨多那() feud.Barony
	BShankargarh商羯罗姞利呬() feud.Barony
	BSharog舍卢俱() feud.Barony
}

type 古卢耆GurgiCounty struct {
	feud.BaseCounty
	婆多迦呼罗Bathgahora   feud.Barony
	旃陀梨醯Chandrehi     feud.Barony
	古卢耆Gurgi          feud.Barony
	沙吞Sathon          feud.Barony
	萨多那Satna          feud.Barony
	商羯罗姞利呬Shankargarh feud.Barony
	舍卢俱Sharog         feud.Barony
}

func (c *古卢耆GurgiCounty) BBathgahora婆多迦呼罗() feud.Barony {
	return c.婆多迦呼罗Bathgahora
}

func (c *古卢耆GurgiCounty) BChandrehi旃陀梨醯() feud.Barony {
	return c.旃陀梨醯Chandrehi
}

func (c *古卢耆GurgiCounty) BGurgi古卢耆() feud.Barony {
	return c.古卢耆Gurgi
}

func (c *古卢耆GurgiCounty) BSathon沙吞() feud.Barony {
	return c.沙吞Sathon
}

func (c *古卢耆GurgiCounty) BSatna萨多那() feud.Barony {
	return c.萨多那Satna
}

func (c *古卢耆GurgiCounty) BShankargarh商羯罗姞利呬() feud.Barony {
	return c.商羯罗姞利呬Shankargarh
}

func (c *古卢耆GurgiCounty) BSharog舍卢俱() feud.Barony {
	return c.舍卢俱Sharog
}

var CGurgi古卢耆 GurgiCounty = &古卢耆GurgiCounty{}

func init() {
	f := CGurgi古卢耆.(*古卢耆GurgiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1278",
		Title:     "gurgi",
		TitleName: "古卢耆",
		TitleCode: "c_gurgi",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆多迦呼罗Bathgahora = BBathgahora婆多迦呼罗
	f.婆多迦呼罗Bathgahora.SetParent(f)

	f.旃陀梨醯Chandrehi = BChandrehi旃陀梨醯
	f.旃陀梨醯Chandrehi.SetParent(f)

	f.古卢耆Gurgi = BGurgi古卢耆
	f.古卢耆Gurgi.SetParent(f)

	f.沙吞Sathon = BSathon沙吞
	f.沙吞Sathon.SetParent(f)

	f.萨多那Satna = BSatna萨多那
	f.萨多那Satna.SetParent(f)

	f.商羯罗姞利呬Shankargarh = BShankargarh商羯罗姞利呬
	f.商羯罗姞利呬Shankargarh.SetParent(f)

	f.舍卢俱Sharog = BSharog舍卢俱
	f.舍卢俱Sharog.SetParent(f)

}
