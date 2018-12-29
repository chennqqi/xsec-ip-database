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

package feeds

import (
	"github.com/chennqqi/xsec-ip-database/models"
	"github.com/chennqqi/xsec-ip-database/util"

	"io/ioutil"
	"strings"
)

func FetchFromFeodotracker() (evilIps models.EvilIps, err error) {
	url := "https://feodotracker.abuse.ch/blocklist/?download=ipblocklist"
	src := "feodotracker.abuse.ch"
	desc := "malware"

	evilIps.Src.Source = src
	evilIps.Src.Desc = desc

	resp, err := util.GetPage(url)
	if err == nil {
		ret, err := ioutil.ReadAll(resp)
		if err == nil {
			lines := strings.Split(string(ret), "\n")
			for _, line := range lines {
				if strings.HasPrefix(line, "#") || !strings.Contains(line, ".") {
					continue
				}
				evilIps.Ips = append(evilIps.Ips, line)
			}
		}
	}
	return evilIps, err
}
