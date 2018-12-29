/*

Copyright (c) 2017 xsec.io

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THEq
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

*/

package settings

import (
	"gopkg.in/ini.v1"

	"github.com/chennqqi/xsec-ip-database/logger"
)

var (
	Cfg    *ini.File
	DEBUG  bool
	SECRET string
)

func init() {
	var err error
	source := "conf/app.ini"
	Cfg, err = ini.Load(source)

	if err != nil {
		logger.Logger.Panicln(err)
	}
	DEBUG = Cfg.Section("").Key("DEBUG").MustBool(true)
	SECRET = Cfg.Section("").Key("SECRET").MustString("SECRET_KEY")
}
