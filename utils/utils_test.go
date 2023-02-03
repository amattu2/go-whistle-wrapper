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

package utils

import (
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGetEnv(t *testing.T) {
	// Set environment variable
	os.Setenv("TEST_ENV_1", "test")

	assert.Equal(t, "test", GetEnv("TEST_ENV_1", ""))
}

func TestGetEnvFallback(t *testing.T) {
	assert.Equal(t, "FB", GetEnv("TEST_ENV_2", "FB"))
}
