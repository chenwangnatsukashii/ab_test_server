package service

import (
	"encoding/json"
	"gorm.io/gorm"
	"line_china/ab_test_server/src/dao"
	"line_china/ab_test_server/src/model"
	"line_china/ab_test_server/src/utils"
	common "line_china/common/model"
	"strconv"
	"time"
)

// AddPublish 添加
func AddPublish(expConfig common.JSON) (error, int) {
	entity := model.PublishConfig{ExpConfig: expConfig, CreateTime: time.Now()}
	return dao.AddPublish(entity)
}

func PublishAll(id int) (map[string]interface{}, error) {
	var err error
	var app *model.Application // 获取App
	app, err = SelectApp(id)
	publishMap := map[string]interface{}{}
	if err == nil {
		publishMap["app_id"] = app.ID
		publishMap["name"] = app.Name

		var domains []model.Domain // 获取域
		domains, err = SelectOnLineDomains(id)
		if err == nil {
			var domainList = make([]interface{}, 0)
			for _, domain := range domains {
				domainMap := map[string]interface{}{}
				domainMap["id"] = domain.ID
				domainMap["name"] = domain.Name
				domainMap["weight"] = domain.Weight

				var layers []model.Layer // 获取层
				layers, err = SelectOnLineLayers(domain.ID)
				if err == nil {
					var layerList = make([]interface{}, 0)
					for _, layer := range layers {
						layerMap := map[string]interface{}{}
						layerMap["id"] = layer.ID
						layerMap["name"] = layer.Name

						var exps []model.Experiment // 获取实验
						exps, err = SelectOnLineExps(layer.ID)
						if err == nil {
							var expList = make([]model.Experiment, 0)
							for _, exp := range exps {
								expList = append(expList, exp)
							}
							layerMap["exps"] = expList
							layerList = append(layerList, layerMap)
						}
						domainMap["layers"] = layerList
					}
				}
				domainList = append(domainList, domainMap)
			}

			publishMap["domains"] = domainList

			var publishOne []byte
			publishOne, err = json.Marshal(publishMap)
			if err == nil {
				err = utils.GormDb.Transaction(func(tx *gorm.DB) error {
					err = OnLineApp(id)
					if err == nil {
						err, id = AddPublish(publishOne)
						if err == nil {
							publishMap["publish_id"] = id

							var publishTwo []byte
							publishTwo, err = json.Marshal(publishMap)
							if err == nil {
								err = utils.PublishConfig(string(publishTwo))
							}
						}
					}
					return err
				})
			}
		}
	}
	if err == nil {
		return publishMap, err
	}
	return nil, err
}

func GetPublishReview(id int) (map[string]interface{}, error) {
	var err error
	var app *model.Application // 获取App
	app, err = SelectApp(id)
	reviewMap := map[string]interface{}{}
	if err == nil {
		reviewMap["id"] = "app" + strconv.Itoa(app.ID)
		reviewMap["label"] = app.Name

		var domains []model.Domain // 获取域
		domains, err = SelectOnLineDomains(id)
		if err == nil {
			var domainList = make([]interface{}, 0)
			for _, domain := range domains {
				domainMap := map[string]interface{}{}
				domainMap["id"] = "domain" + strconv.Itoa(domain.ID)
				domainMap["label"] = domain.Name
				domainMap["weight"] = domain.Weight

				var layers []model.Layer // 获取层
				layers, err = SelectOnLineLayers(domain.ID)
				if err == nil {
					var layerList = make([]interface{}, 0)
					for _, layer := range layers {
						layerMap := map[string]interface{}{}
						layerMap["id"] = "layer" + strconv.Itoa(layer.ID)
						layerMap["label"] = layer.Name

						var exps []model.Experiment // 获取实验
						exps, err = SelectOnLineExps(layer.ID)
						if err == nil {
							var expList = make([]interface{}, 0)
							for _, exp := range exps {
								var oneExp = map[string]interface{}{}
								oneExp["id"] = "exp" + strconv.Itoa(exp.ID)
								oneExp["label"] = exp.Name
								oneExp["weight"] = exp.Weight
								expList = append(expList, oneExp)
							}
							layerMap["children"] = expList
							layerList = append(layerList, layerMap)
						}
					}
					domainMap["children"] = layerList
					domainList = append(domainList, domainMap)
				}
			}
			reviewMap["children"] = domainList
		}
	}
	if err == nil {
		return reviewMap, err
	}
	return nil, err
}
