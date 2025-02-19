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
package forexCore

/*
 * The function takes in selling and buying rates and uses
 * this formula to calculate it.
 *     average rate = selling rate + buying rate
 *                    --------------------------
 *                                 2
 */
func CalculateAveragegRate(selling_rate, buy_rate float64) (average_rate float64) {
	if selling_rate <= 0 || buy_rate <= 0 {
		return 0
	}
	return (selling_rate + buy_rate) / 2
}
