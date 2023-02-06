/*
 * Produced: Fri Feb 03 2023
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

package whistle_test

import (
	"testing"

	"github.com/amattu2/go-whistle-wrapper/utils"
	"github.com/amattu2/go-whistle-wrapper/whistle"
	"github.com/go-playground/assert/v2"
)

var (
	Email    = utils.GetEnv("EMAIL", "")
	Password = utils.GetEnv("PASSWORD", "")
)

func TestInvalidInit(t *testing.T) {
	t.Parallel()

	assert.PanicMatches(t, func() {
		whistle.Initialize("", "")
	}, "valid email and password are required")
}

func TestInvalidBearerInit(t *testing.T) {
	t.Parallel()

	assert.PanicMatches(t, func() {
		whistle.InitializeBearer("")
	}, "valid http bearer is required")
}

func TestInvalidTokenInit(t *testing.T) {
	t.Parallel()

	assert.PanicMatches(t, func() {
		whistle.InitializeToken("")
	}, "valid API token is required")
}

func TestInvalidRefreshInit(t *testing.T) {
	t.Parallel()

	assert.PanicMatches(t, func() {
		whistle.InitializeRefreshToken("abc@gmail.com", "")
	}, "valid email and refresh token are required")
}
