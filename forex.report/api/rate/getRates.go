/**
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package forexReport

import (
	"encoding/json"
	"time"

	"api.forex.com/forex.report/common"
	"api.forex.com/models"
	"api.forex.com/storage"
	"api.forex.com/utilities"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

func GetRate(ctx iris.Context) {
	//check if the user performing this operation has adequate permissions.
	var input common.GetRateInput
	body, err := ctx.GetBody()
	if err != nil {
		utilities.CreateInternalServerError(ctx)
		return
	}
	if len(body) == 0 {
		//set current date as a default date values if not passed in.
		input.From = time.Now().Format("2006-01-02")
		input.To = time.Now().Format("2006-01-02")
	} else {
		var requestData common.GetRateInput
		if err := json.Unmarshal(body, &requestData); err != nil {
			utilities.CreateError(iris.StatusBadRequest, "Invalid json format", "Invalid json format", ctx)
			return
		}
		input.From = requestData.From
		input.To = requestData.To
		if requestData.CurrencyID != uuid.Nil {
			input.CurrencyID = requestData.CurrencyID
		}
	}
	//query the database for the requested data
	var rates []models.Rate
	query := storage.DB.Where("created_date BETWEEN ? AND ?", input.From, input.To).
		Preload("FkCurrencyID").
		Preload("FkCreatedbyID").
		Order("created_date, created_time DESC")
	if input.CurrencyID != uuid.Nil {
		query = query.Where("currency_id IN (?)", input.CurrencyID)
	}
	query.Find(&rates)
	ctx.JSON(iris.Map{
		"totalCount": len(rates),
		"rates":      rates,
	})
}
