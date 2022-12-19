// Copyright 2016 NDP Systèmes. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package base

import (
	"github.com/hexya-erp/hexya/src/models"
	data "github.com/hexya-addons/base/models"
	"github.com/hexya-erp/hexya/src/models/security"
	"github.com/hexya-erp/hexya/src/server"
	"github.com/hexya-erp/hexya/src/tools/logging"
)

const (
	// MODULE_NAME is the name of this module
	MODULE_NAME = "base"
)

var log logging.Logger

// Call each model to register and initialize in registry
func initModels() {
	log.Info("Loading settings..", MODULE_NAME)
	data.InitSettings()
	log.Info("Loading attachments..", MODULE_NAME)
	data.InitAttachments()
	log.Info("Loading bank information..", MODULE_NAME)
	data.InitBank()
	log.Info("Loading lang and translations..", MODULE_NAME)
	data.InitTranslations()
	log.Info("Loading sequences..", MODULE_NAME)
	data.InitSequences()
	log.Info("Loading jobs and queue channels..", MODULE_NAME)
	data.InitJobQueues()
	log.Info("Loading users..", MODULE_NAME)
	data.InitUsers()
	log.Info("Loaded..", MODULE_NAME)
}

func init() {
	log = logging.GetLogger("base")
	server.RegisterModule(&server.Module{
		Name:    MODULE_NAME,
		PreInit: initModels,
		PostInit: func() {
			err := models.ExecuteInNewEnvironment(security.SuperUserID, func(env models.Environment) {
				//h.Group().NewSet(env).ReloadGroups()
			})
			if err != nil {
				log.Panic("Error while initializing", "error", err)
			}
		},
	})
}
