// Copyright 2019 Google LLC
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

// Code generated by "mdtogo"; DO NOT EDIT.
package overview

var ReferenceShort = `Overview of kpt commands`
var ReferenceLong = `
kpt functionality is subdivided into the following command groups, each of
which operates on a particular set of entities, with a consistent command
syntax and pattern of inputs and outputs.

| Command Group | Description                                                                     | Reads From      | Writes To       |
| ------------- | ------------------------------------------------------------------------------- | --------------- | --------------- |
| [pkg]         | fetch, update, and sync configuration files using git                           | remote git      | local directory |
| [cfg]         | examine and modify configuration files                                          | local directory | local directory |
| [fn]          | generate, transform, validate configuration files using containerized functions | local directory | local directory |
| [live]        | reconcile the live state with configuration files                               | local directory | remote cluster  |
`
var ReferenceExamples = `
  # get a package
  $ kpt pkg get https://github.com/GoogleContainerTools/kpt.git/package-examples/helloworld-set@v0.5.0 helloworld
  fetching package /package-examples/helloworld-set from \
    https://github.com/GoogleContainerTools/kpt to helloworld

  # list setters and set a value
  $ kpt cfg list-setters helloworld
  NAME            DESCRIPTION         VALUE    TYPE     COUNT   SETBY
  http-port   'helloworld port'         80      integer   3
  image-tag   'hello-world image tag'   v0.3.0  string    1
  replicas    'helloworld replicas'     5       integer   1
  
  $ kpt cfg set helloworld replicas 3 --set-by pwittrock  --description 'reason'
  set 1 fields

  # get a package and run a validation function
  kpt pkg get https://github.com/GoogleContainerTools/kpt-functions-sdk.git/example-configs example-configs
  mkdir results/
  kpt fn run example-configs/ --results-dir results/ --image gcr.io/kpt-functions/validate-rolebinding:results -- subject_name=bob@foo-corp.com

  # apply the package to a cluster
  $ kpt live apply --reconcile-timeout=10m helloworld
  ...
  all resources has reached the Current status
`
