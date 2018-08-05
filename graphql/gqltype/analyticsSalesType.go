package gqltype

import (
	"github.com/graphql-go/graphql"
	"colorme.vn/model"
	"colorme.vn/core/service"
	"time"
)

type analyticSales struct {
	Money      int
	MoneyToday int
}

var RegisterByDate = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RegisterByDate",
		Fields: graphql.Fields{
			"date": &graphql.Field{
				Type: graphql.String,
			},
			"total": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var AnalyticSalesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AnalyticSalesType",
		Fields: graphql.Fields{
			"money": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args := p.Source.(map[string]interface{})

					genID, ok := args["gen_id"].(int)
					if !ok {
						genID = int(model.GetCurrentGen().ID)
					}

					db := service.GetService().DB.DB

					var gen model.Gen
					var sales analyticSales

					db.Find(&gen, genID)

					db.Model(&model.Register{}).Select("sum(money) as money").
						Where("? <= paid_time AND paid_time < ?", gen.StartTime, gen.EndTime.AddDate(0, 0, 1)).
						Scopes(model.PaidMoney).Scan(&sales)

					return sales.Money, nil
				},
			},
			"money_today": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					db := service.GetService().DB.DB

					var sales analyticSales

					dateNow := time.Now()

					db.Model(&model.Register{}).Select("sum(money) as money_today").
						Where("DATE(?) = DATE(paid_time)", dateNow).
						Scopes(model.PaidMoney).Scan(&sales)

					return sales.MoneyToday, nil
				},
			},
			"registers_today": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					db := service.GetService().DB.DB

					dateNow := time.Now()

					var registersToday int

					db.Model(&model.Register{}).
						Where("DATE(?) = DATE(created_at)", dateNow).
						Scopes(model.RegisterNew).
						Count(&registersToday)

					return registersToday, nil
				},
			},
			"total_paid_registers": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					args := p.Source.(map[string]interface{})

					genID, ok := args["gen_id"].(int)
					if !ok {
						genID = int(model.GetCurrentGen().ID)
					}

					db := service.GetService().DB.DB

					var gen model.Gen

					var total int

					db.Find(&gen, genID)

					db.Model(&model.Register{}).
						Where("? <= paid_time AND paid_time < ?", gen.StartTime, gen.EndTime.AddDate(0, 0, 1)).
						Scopes(model.PaidMoney).
						Count(&total)

					return total, nil
				},
			},
			"registers_by_date": &graphql.Field{
				Type: graphql.NewList(RegisterByDate),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args := p.Source.(map[string]interface{})

					genID, ok := args["gen_id"].(int)
					if !ok {
						genID = int(model.GetCurrentGen().ID)
					}

					db := service.GetService().DB.DB

					var gen model.Gen

					type result struct {
						DateTime time.Time
						Date     string `json:"date"`
						Total    int    `json:"total"`
					}

					var results []result

					db.Find(&gen, genID)

					db.Debug().Model(&model.Register{}).
						Select("date(created_at) date_time, sum(1) as total").
						Where("? <= created_at AND created_at < ?", gen.StartTime, gen.EndTime.AddDate(0, 0, 1)).
						Scopes(model.RegisterNew).
						Group("date(created_at)").Scan(&results)

					for i := 0; i < len(results); i++ {
						results[i].Date = results[i].DateTime.Format("02-01-2006")
					}

					return results, nil
				},
			},
			"paid_by_date": &graphql.Field{
				Type: graphql.NewList(RegisterByDate),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args := p.Source.(map[string]interface{})

					genID, ok := args["gen_id"].(int)
					if !ok {
						genID = int(model.GetCurrentGen().ID)
					}

					db := service.GetService().DB.DB

					var gen model.Gen

					type result struct {
						DateTime time.Time
						Date     string `json:"date"`
						Total    int    `json:"total"`
					}

					var results []result

					db.Find(&gen, genID)

					db.Model(&model.Register{}).
						Select("date(paid_time) date_time,sum(1) as total").
						Where("? <= paid_time AND paid_time < ?", gen.StartTime, gen.EndTime.AddDate(0, 0, 1)).
						Scopes(model.PaidMoney).
						Group("date(paid_time)").Scan(&results)

					for i := 0; i < len(results); i++ {
						results[i].Date = results[i].DateTime.Format("02-01-2006")
					}

					return results, nil
				},
			},
			"target_revenue": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args := p.Source.(map[string]interface{})

					genID, ok := args["gen_id"].(int)
					if !ok {
						genID = int(model.GetCurrentGen().ID)
					}

					db := service.GetService().DB.DB

					var gen model.Gen

					db.Find(&gen, genID)

					db.Model(&gen).Related(&gen.Classes)

					targetRevenue := model.GetTargetClasses(&gen.Classes)

					return targetRevenue, nil
				},
			},
		},
	},
)
