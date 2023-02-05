/*
 * Produced: Sat Feb 04 2023
 * Author: Alec M.
 * GitHub: https://amattu.com/links/github
 * Copyright: (C) 2023 Alec M.
 * License: License GNU Affero General Public License v3.0
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package whistle

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type BreedsResponse struct {
	Breeds []Breed `json:"breeds"`
}

type Breed struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Popularity int    `json:"popularity"`
}

// Breeds returns a list of breeds for a given animal (dogs, cats)
func (c Client) Breeds(animal string) *HttpResponse[BreedsResponse] {
	resp, err := c.get(fmt.Sprintf("api/breeds/%s", animal), nil, true)
	if err != nil || resp.StatusCode != http.StatusOK {
		return &HttpResponse[BreedsResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
			Raw:        resp,
		}
	}

	defer resp.Body.Close()

	// Parse json response
	body, _ := io.ReadAll(resp.Body)
	result := BreedsResponse{}
	json.Unmarshal(body, &result)

	return &HttpResponse[BreedsResponse]{
		StatusCode: resp.StatusCode,
		Response:   result,
		Raw:        resp,
	}
}
