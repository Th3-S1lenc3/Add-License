package main

// isNotExist.go
//
// part of Add-License
//
// Copyright (C) 2022 Th3-S1lenc3
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

import "os"

// Returns true if path does not exist
func IsNotExist(path string) bool {
  _, err := os.Stat(path)
  if err != nil && os.IsNotExist(err) {
    return true
  }

  return false
}
