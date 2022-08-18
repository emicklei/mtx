package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/emicklei/moneypenny/aws"
	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/gcp/bq"
)

func main() {
	data, err := os.ReadFile("/Users/emicklei/xProjects/moneypenny/aws/aws-billing-report-Manifest.json")
	if err != nil {
		panic(err)
	}
	m := new(aws.BillingReportManifest)
	err = json.Unmarshal(data, m)
	if err != nil {
		panic(err)
	}

	db := bq.NewDataset("aws-cur")
	tab := db.Table("aws-billing-reports")
	for _, each := range m.Columns {
		dt, canRepresentNull := mapType(each.Type)
		col := tab.C(fmt.Sprintf("%s_%s", each.Category, each.Name), dt, fmt.Sprintf("category:%s,name:%s,type:%s", each.Category, each.Name, each.Type))
		col.IsNotNull = !canRepresentNull

	}
	fmt.Println(bq.ToJSONSchema(tab))
}

func mapType(name string) (mtx.Datatype, bool) {
	switch name {
	case "String":
		return bq.STRING, false
	case "OptionalString":
		return bq.STRING, true
	case "Interval":
		return bq.INTERVAL, false
	case "DateTime":
		return bq.DATETIME, false
	case "BigDecimal":
		return bq.BigDecimal(15, 5), false
	case "OptionalBigDecimal":
		return bq.BigDecimal(15, 5), true
	default:
		fmt.Println("WARN:", name)
		return bq.STRING, false
	}
}
