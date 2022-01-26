// Copyright Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ossm

import (
	"strings"
	"testing"
	"time"

	"github.com/maistra/maistra-test-tool/pkg/examples"
	"github.com/maistra/maistra-test-tool/pkg/util"
)

func cleanupBookinfo() {
	util.Log.Info("Cleanup")
	app := examples.Bookinfo{"bookinfo"}
	app.Uninstall()
	time.Sleep(time.Duration(30) * time.Second)
}

func TestBookinfo(t *testing.T) {
	defer util.RecoverPanic(t)

	util.Log.Info("Test Bookinfo Installation")
	app := examples.Bookinfo{"bookinfo"}
	app.Install(false)

	util.Log.Info("Check pods running 2/2 ready")
	msg, _ := util.Shell(`oc get pods -n bookinfo`)
	if strings.Contains(msg, "2/2") {
		util.Log.Info("Success. proxy container is running.")
	} else {
		t.Error("Error. proxy container is not running.")
	}
}