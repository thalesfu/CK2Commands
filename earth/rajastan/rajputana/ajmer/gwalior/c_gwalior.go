package gwalior

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GwaliorCounty interface {
	feud.County
	BBateshwar婆知湿伐罗() feud.Barony
	BDabra德布拉() feud.Barony
	BGwalior伽婆梨耶罗() feud.Barony
	BNarwar讷尔沃尔() feud.Barony
	BSabalgarh娑婆罗姞利呬() feud.Barony
	BSahastrabahu娑诃斯罗婆呼() feud.Barony
	BVijaypur毗阇耶补罗() feud.Barony
}

type 伽婆梨耶罗GwaliorCounty struct {
	feud.BaseCounty
	婆知湿伐罗Bateshwar     feud.Barony
	德布拉Dabra           feud.Barony
	伽婆梨耶罗Gwalior       feud.Barony
	讷尔沃尔Narwar         feud.Barony
	娑婆罗姞利呬Sabalgarh    feud.Barony
	娑诃斯罗婆呼Sahastrabahu feud.Barony
	毗阇耶补罗Vijaypur      feud.Barony
}

func (c *伽婆梨耶罗GwaliorCounty) BBateshwar婆知湿伐罗() feud.Barony {
	return c.婆知湿伐罗Bateshwar
}

func (c *伽婆梨耶罗GwaliorCounty) BDabra德布拉() feud.Barony {
	return c.德布拉Dabra
}

func (c *伽婆梨耶罗GwaliorCounty) BGwalior伽婆梨耶罗() feud.Barony {
	return c.伽婆梨耶罗Gwalior
}

func (c *伽婆梨耶罗GwaliorCounty) BNarwar讷尔沃尔() feud.Barony {
	return c.讷尔沃尔Narwar
}

func (c *伽婆梨耶罗GwaliorCounty) BSabalgarh娑婆罗姞利呬() feud.Barony {
	return c.娑婆罗姞利呬Sabalgarh
}

func (c *伽婆梨耶罗GwaliorCounty) BSahastrabahu娑诃斯罗婆呼() feud.Barony {
	return c.娑诃斯罗婆呼Sahastrabahu
}

func (c *伽婆梨耶罗GwaliorCounty) BVijaypur毗阇耶补罗() feud.Barony {
	return c.毗阇耶补罗Vijaypur
}

var CGwalior伽婆梨耶罗 GwaliorCounty = &伽婆梨耶罗GwaliorCounty{}

func init() {
	f := CGwalior伽婆梨耶罗.(*伽婆梨耶罗GwaliorCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1172",
		Title:     "gwalior",
		TitleName: "伽婆梨耶罗",
		TitleCode: "c_gwalior",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆知湿伐罗Bateshwar = BBateshwar婆知湿伐罗
	f.婆知湿伐罗Bateshwar.SetParent(f)

	f.德布拉Dabra = BDabra德布拉
	f.德布拉Dabra.SetParent(f)

	f.伽婆梨耶罗Gwalior = BGwalior伽婆梨耶罗
	f.伽婆梨耶罗Gwalior.SetParent(f)

	f.讷尔沃尔Narwar = BNarwar讷尔沃尔
	f.讷尔沃尔Narwar.SetParent(f)

	f.娑婆罗姞利呬Sabalgarh = BSabalgarh娑婆罗姞利呬
	f.娑婆罗姞利呬Sabalgarh.SetParent(f)

	f.娑诃斯罗婆呼Sahastrabahu = BSahastrabahu娑诃斯罗婆呼
	f.娑诃斯罗婆呼Sahastrabahu.SetParent(f)

	f.毗阇耶补罗Vijaypur = BVijaypur毗阇耶补罗
	f.毗阇耶补罗Vijaypur.SetParent(f)

}
