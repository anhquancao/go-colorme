package gqltype

import (
	"github.com/graphql-go/graphql"
	"fmt"
	"colorme.vn/model"
	"colorme.vn/core/service"
	"time"
)

type analyticSales struct {
	Money      int
	MoneyToday int
}

var AnalyticSalesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Register",
		Fields: graphql.Fields{
			"money": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args := p.Source.(map[string]interface{})

					genID, ok := args["gen_id"].(int)
					fmt.Println(genID)
					if !ok {
						genID = int(model.GetCurrentGen().ID)
					}

					db := service.GetService().DB.DB

					var gen model.Gen
					var sales analyticSales

					db.Find(&gen, genID)

					db.Table("registers").Select("sum(money) as money").
						Where("? < paid_time AND paid_time < ?", gen.StartTime, gen.EndTime).
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

					db.Table("registers").Select("sum(money) as money_today").
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

					db.Table("registers").
						Where("DATE(?) = DATE(created_at)", dateNow).
						Scopes(model.RegisterNew).
						Count(&registersToday)

					return registersToday, nil
				},
			},
			"total_paid_registers": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					db := service.GetService().DB.DB

					dateNow := time.Now()

					var registersToday int

					db.Table("registers").
						Where("DATE(?) = DATE(paid_time)", dateNow).
						Scopes(model.PaidMoney).
						Count(&registersToday)

					return registersToday, nil
				},
			},
			"registers_by_date": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args := p.Source.(map[string]interface{})

					genID, ok := args["gen_id"].(int)
					fmt.Println(genID)
					if !ok {
						genID = int(model.GetCurrentGen().ID)
					}

					db := service.GetService().DB.DB

					var gen model.Gen

					register := model.Register{}

					var results []int

					db.Find(&gen, genID)

					db.Debug().Table(register.TableName()).
						Select("sum(1) as total").
						Where("? < created_at AND created_at < ?", gen.StartTime, gen.EndTime).
						Scopes(model.RegisterNew).
						Group("date(created_at)").Pluck("total", &results)

					fmt.Println(results)

					return results, nil
				},
			},
			"paid_by_date": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args := p.Source.(map[string]interface{})

					genID, ok := args["gen_id"].(int)
					fmt.Println(genID)
					if !ok {
						genID = int(model.GetCurrentGen().ID)
					}

					db := service.GetService().DB.DB

					var gen model.Gen

					register := model.Register{}

					var results []int

					db.Find(&gen, genID)

					db.Debug().Table(register.TableName()).
						Select("sum(1) as total").
						Where("? < created_at AND created_at < ?", gen.StartTime, gen.EndTime).
						Scopes(model.PaidMoney).
						Group("date(created_at)").Pluck("total", &results)

					fmt.Println(results)

					return results, nil
				},
			},
		},
	},
)
